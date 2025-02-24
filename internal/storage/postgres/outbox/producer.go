package outbox

import (
	"errors"
	"log/slog"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/wjojf/go-uber-fx/internal/events"
)

var (
	p *Producer = nil

	ErrProducerAlreadyStarted = errors.New("producer already started")
)

type Producer struct {
	log     *slog.Logger
	pool    *pgxpool.Pool
	started bool

	publisher    events.Publisher
	ticker       *time.Ticker
	deleteTicker *time.Ticker

	doneCh chan struct{}
}

func NewProducer(log *slog.Logger, pool *pgxpool.Pool, publisher events.Publisher) *Producer {

	if p != nil {
		return p
	}

	return &Producer{
		log:          log,
		pool:         pool,
		publisher:    publisher,
		started:      false,
		ticker:       time.NewTicker(500 * time.Millisecond),
		deleteTicker: time.NewTicker(5 * time.Second),
		doneCh:       make(chan struct{}),
	}
}

func (p *Producer) Stop() {
	if !p.started {
		return
	}

	p.started = false
	p.ticker.Stop()
	p.deleteTicker.Stop()

	p.doneCh <- struct{}{}

	close(p.doneCh)
}

func (p *Producer) StartProducing() error {
	if p.started {
		return ErrProducerAlreadyStarted
	}

	p.started = true
	return p.listen()
}

func (p *Producer) listen() error {

	p.log.Info("outbox producer started")

	for {
		select {
		case <-p.ticker.C:
			//p.log.Info("checking for outbox events to produce")

			count, err := p.produce()
			if err != nil {
				p.log.Error("failed to produce outbox events", slog.Any("err", err))
				continue
			}
			if count != 0 {
				p.log.Info("outbox events produced", slog.Int("count", count))
			}

		case <-p.deleteTicker.C:
			p.log.Debug("checking for outbox events to delete")

			err := DeleteProducesOutboxEvents(p.pool)
			if err != nil {
				p.log.Error("failed to delete produced outbox events", slog.Any("err", err))
				continue
			}

			p.log.Debug("produced outbox events deleted")
		case <-p.doneCh:
			p.log.Info("Received stop signal. outbox producer stopped")
			break
		}
	}
}

func (p *Producer) produce() (int, error) {
	messages, err := GetOutboxEventsToProduce(p.pool)
	if err != nil {
		return 0, err
	}

	for _, message := range messages {
		err := p.publisher.Publish(message.EventName, message.Payload)
		if err != nil {
			return 0, err
		}

		err = MarkOutboxEventAsProduced(p.pool, message.ID)
		if err != nil {
			p.log.Error("failed to mark outbox event as produced", slog.Any("err", err))
			return 0, err
		}

	}

	return len(messages), nil
}

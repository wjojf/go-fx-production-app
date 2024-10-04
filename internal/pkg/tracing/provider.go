package tracing

import (
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegerCfg "github.com/uber/jaeger-client-go/config"
	"github.com/wjojf/go-uber-fx/internal/pkg/config"
	"io"
)

func New(cfg config.Config) (opentracing.Tracer, io.Closer, error) {
	c := jaegerCfg.Configuration{
		ServiceName: "golang-uber-fx-backend",
		Sampler: &jaegerCfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1, // sample all traces
		},
		Reporter: &jaegerCfg.ReporterConfig{
			LocalAgentHostPort: cfg.JaegerUrl,
			LogSpans:           true,
		},
	}

	tracer, closer, err := c.NewTracer()
	if err != nil {
		return nil, nil, err
	}

	// Set the tracer as the global tracer
	opentracing.SetGlobalTracer(tracer)

	return tracer, closer, nil

}

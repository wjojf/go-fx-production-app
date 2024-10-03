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
			Type:  jaeger.SamplerTypeRateLimiting,
			Param: 100, // 100 traces per second
		},
		Reporter: &jaegerCfg.ReporterConfig{
			LocalAgentHostPort: cfg.JaegerUrl,
			LogSpans:           true,
		},
	}

	return c.NewTracer()
}

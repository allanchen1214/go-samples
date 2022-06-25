package main

import (
	"fmt"
	"io"
	"os"

	opentracing "github.com/opentracing/opentracing-go"
	log "github.com/opentracing/opentracing-go/log"
	jaeger "github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
)

// initJaeger returns an instance of Jaeger Tracer that samples 100% of traces and logs all spans to stdout.
func initJaeger(service string) (opentracing.Tracer, io.Closer) {
	cfg := &jaegercfg.Configuration{
		ServiceName: service,
		Sampler: &jaegercfg.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: "127.0.0.1:6831",
		},
	}
	tracer, closer, err := cfg.NewTracer(jaegercfg.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	return tracer, closer
}

func formatString(parentSpan opentracing.Span, helloTo string) string {
	span := parentSpan.Tracer().StartSpan("formatString", opentracing.ChildOf(parentSpan.Context()))
	defer span.Finish()

	helloStr := innerFormatString(span, helloTo)
	span.LogFields(
		log.String("event", "string-format"),
		log.String("value", helloStr),
	)

	return helloStr
}

func innerFormatString(parentSpan opentracing.Span, helloTo string) string {
	span := parentSpan.Tracer().StartSpan("innerFormatString", opentracing.ChildOf(parentSpan.Context()))
	defer span.Finish()

	helloStr := fmt.Sprintf("Hello, %s!", helloTo)
	span.LogFields(
		log.String("event", "[inner func]string-format"),
		log.String("value", helloStr),
	)

	return helloStr
}

func printHello(parentSpan opentracing.Span, helloStr string) {
	span := parentSpan.Tracer().StartSpan("printHello", opentracing.ChildOf(parentSpan.Context()))
	defer span.Finish()

	println(helloStr)
	span.LogKV("event", "println")
}

func main() {
	if len(os.Args) != 2 {
		panic("ERROR: Expecting one argument")
	}

	tracer, closer := initJaeger("hello-world")
	defer closer.Close()

	helloTo := os.Args[1]

	span := tracer.StartSpan("say-hello")
	span.SetTag("hello-to", helloTo)
	defer span.Finish()

	helloStr := formatString(span, helloTo)
	printHello(span, helloStr)
}

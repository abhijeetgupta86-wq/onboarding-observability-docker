package main

import (
	"context"
	"log"
	"os"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

var tracer = otel.Tracer("onboarding-demo")

func initTelemetry() {
	endpoint := os.Getenv("DT_ENDPOINT")
	token := os.Getenv("DT_TOKEN")

	exporter, err := otlptracehttp.New(
		context.Background(),
		otlptracehttp.WithEndpoint(endpoint),
		otlptracehttp.WithHeaders(map[string]string{
			"Authorization": "Api-Token " + token,
		}),
	)

	if err != nil {
		log.Fatal(err)
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
	)

	otel.SetTracerProvider(tp)
	log.Println("Telemetry initialized")
}

func startSpan(name string) (context.Context, sdktrace.Span) {
	return tracer.Start(context.Background(), name)
}

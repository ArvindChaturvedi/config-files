package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/signalfx/splunk-otel-go/distro"
	"go.opentelemetry.io/otel/exporters/otlp/otlphttp"
	"go.opentelemetry.io/otel/trace"
)

func main() {
	// Create a new OTLP exporter to send the tracing data to the Splunk OTel Collector
	exporter, err := otlphttp.New(context.Background(),
		otlphttp.WithEndpoint(os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT")),
	)
	if err != nil {
		fmt.Println("Error creating OTLP exporter:", err)
		return
	}

	// Create a new SimpleSpanProcessor and set the exporter as its receiver
	ssp := distro.NewSimpleSpanProcessor(exporter)
	tp := distro.NewTracerProvider(distro.WithSpanProcessor(ssp))

	// Set the global trace provider
	trace.SetTracerProvider(tp)

	// Make an HTTP request
	resp, err := http.Get("https://example.com")
	if err != nil {
		fmt.Println("Error making HTTP request:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("HTTP request successful")
}
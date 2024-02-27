package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/signalfx/splunk-otel-go/distro"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

func main() {
	// Initialize the Splunk distribution of OpenTelemetry for Go
	distro.Start()

	// Create an HTTP client to make a request
	client := &http.Client{}

	// Create a new request
	req, err := http.NewRequest("GET", "https://example.com/api", nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Create a context to capture the traces for this request
	ctx := context.Background()

	// Make the HTTP request
	resp, err := client.Do(req.WithContext(ctx))
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	// Print the response status code
	fmt.Println("Response Status:", resp.Status)
}
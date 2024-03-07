package main

import (
	"net/http"

	"github.com/signalfx/splunk-otel-go/distro"
	"github.com/signalfx/splunk-otel-go/instrumentation/net/http/splunkhttp"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func main() {
	distro.Run()

	var handler http.Handler = http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello"))
		},
	)
	handler = splunkhttp.NewHandler(handler)
	handler = otelhttp.NewHandler(handler, "my-service")

	http.ListenAndServe(":9090", handler)
}

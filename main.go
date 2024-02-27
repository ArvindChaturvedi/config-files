package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
	"github.com/signalfx/splunk-otel-go/distro"
)

func main() {
	sdk, err := distro.Run()
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := sdk.Shutdown(context.Background()); err != nil {
			panic(err)
		}
	}()

	for {
		resp, err := http.Get("http://www.google.com")
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			resp.Body.Close()
			if resp.StatusCode == http.StatusOK {
				fmt.Println("Google is up and running!")
			} else {
				fmt.Println("Google is not responding correctly")
			}
		}
		time.Sleep(5 * time.Second) // Wait for 5 seconds before pinging again
	}
}
package main

import (
	"context"
	"net/http"
	"os"
	"time"

	"search-service/app/initializer"
)

func main() {
	if len(os.Args) == 2 && os.Args[1] == "healthcheck" {
		if !isHealthy("http://127.0.0.1:9104/metrics") {
			os.Exit(1)
		}
		return
	}

	ps := initializer.InitializePrometheusServer()
	initializer.InitializerConfig()
	initializer.InitializeLogger()
	initializer.InitializeDb()
	go func() {
		initializer.InitializeProductConsumer()
	}()
	initializer.InitializeServer(context.Background(), ps)
}

func isHealthy(endpoint string) bool {
	client := http.Client{Timeout: 2 * time.Second}
	response, err := client.Get(endpoint)
	if err != nil {
		return false
	}
	statusOK := response.StatusCode == http.StatusOK
	if err := response.Body.Close(); err != nil {
		return false
	}
	return statusOK
}

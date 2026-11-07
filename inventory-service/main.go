package main

import (
	"context"
	"net/http"
	"os"
	"time"

	"inventory-service/app/initializer"
	"inventory-service/pkg/logger"
)

func main() {
	if len(os.Args) == 2 && os.Args[1] == "healthcheck" {
		if !isHealthy("http://127.0.0.1:3003/health") {
			os.Exit(1)
		}
		return
	}

	ctx := context.Background()
	initializer.InitializeConfig()
	initializer.InitializeLogger()
	if err := initializer.InitializeDb(); err != nil {
		logger.Error("error in initializing PostgreSQL connection", "error", err)
		panic(err)
	}
	if err := initializer.InitializeServer(ctx); err != nil {
		logger.Error("error in starting http server", "error", err)
		panic(err)
	}
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

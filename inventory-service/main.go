package main

import (
	"context"
	"inventory-service/app/initializer"
	"inventory-service/pkg/logger"
)

func main() {
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

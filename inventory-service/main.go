package main

import (
	"context"
	"inventory-service/app/initializer"
)

func main() {
	ctx := context.Background()
	initializer.InitializeConfig()
	initializer.InitializeLogger()
	initializer.InitializeDb()
	initializer.InitializeServer(ctx)
}

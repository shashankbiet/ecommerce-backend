package main

import (
	"context"
	"search-service/app/initializer"
)

func main() {
	ps := initializer.InitializePrometheusServer()
	initializer.InitializerConfig()
	initializer.InitializeLogger()
	initializer.InitializeDb()
	initializer.InitializeProductConsumer()
	initializer.InitializeServer(context.Background(), ps)
}

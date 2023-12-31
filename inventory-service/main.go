package main

import (
	"inventory-service/app/initializer"
)

func main() {
	initializer.InitializeConfig()
	initializer.InitializeLogger()
	initializer.InitializeDb()
	initializer.InitializeHttp()
}

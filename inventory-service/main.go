package main

import (
	"inventory-service/app/initializer"
)

func main() {
	initializer.InitializeConfig()
	initializer.InitializeDb()
	initializer.InitializeHttp()
}

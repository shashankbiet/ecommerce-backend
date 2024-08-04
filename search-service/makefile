.DEFAULT_GOAL := run

build:
	@go build -o ./.bin/app main.go

run: build
	@./.bin/app

clean:
	go clean
	rm ./.bin/app
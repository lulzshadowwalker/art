.PHONY: build run
APP_NAME=art

run: build
	@echo "Running..."
	./bin/$(APP_NAME) model 

build:
	@echo "Building..."
	@go build -o bin/$(APP_NAME) cmd/cli/main.go
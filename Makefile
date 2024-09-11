.PHONY: run

run:
	@echo "Loading environment variables..."
	@export $(cat .env | xargs) && go run main.go

build:
	@echo "Building the application..."
	@go build -o main .

install-deps:
	@echo "Installing dependencies..."
	@go mod tidy

clean:
	@echo "Cleaning up..."
	@go clean
	@rm -f main
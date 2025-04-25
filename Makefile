run:
	@echo "Running the program..."
	go run ./examples/main.go
	@echo "Program completed."

test:
	@echo "Running tests..."
	go test -v -cover -coverprofile=coverage.out ./manhattan 
	@echo "Tests completed."

bench:
	@echo "Running test bench..."
	go test -bench=. -benchmem ./manhattan
	@echo "Test bench completed."

build:
	@echo "Building manhattan binary..."
	cd ./examples && \
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ./manhattan-app main.go
	@echo "Binary built: manhattan-app"

up:
	@echo Starting Docker images...
	docker-compose up -d
	@echo Docker images started!

down:
	@echo Stopping docker compose...
	docker-compose down
	@echo Done!

up_build: build up
	@echo "Building manhattan binary and starting Docker images..."

phony: run test bench build up down up_build 
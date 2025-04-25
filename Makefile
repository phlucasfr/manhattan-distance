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

phony: run test bench

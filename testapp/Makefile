# on windows add .exe to binaryname
BINARY_NAME=bourbonApp

build:
	@go mod vendor
	@echo "Building Bourbon..."
	@go build -o tmp/${BINARY_NAME} .
	@echo "Bourbon built!"

run: build
	@echo "Starting Bourbon..."
	@./tmp/${BINARY_NAME} &
	@echo "Bourbon started!"

clean:
	@echo "Cleaning..."
	@go clean
	@rm tmp/${BINARY_NAME}
	@echo "Cleaned!"

test:
	@echo "Testing..."
	@go test ./...
	@echo "Done!"

start: run

stop:
	@echo "Stopping Bourbon..."
	@-pkill -SIGTERM -f "./tmp/${BINARY_NAME}"
	@echo "Stopped Bourbon!"

restart: stop start
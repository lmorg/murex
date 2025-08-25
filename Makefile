# Variables
BINARY_NAME=murex
GOFLAGS=-v
BUILD_DIR=./bin
SOURCE_DIR=.

# Build variables that can be overridden
BRANCH?=$(shell git rev-parse --abbrev-ref HEAD || echo "unknown")
BUILD_DATE?=$(shell date -u '+%Y-%m-%d_%H:%M:%S' || echo "unknown")
LDFLAGS=-ldflags "-X github.com/lmorg/murex/app.branch=${BRANCH} -X github.com/lmorg/murex/app.buildDate=${BUILD_DATE}"

# Default target
.PHONY: all
all: build

# Build the binary
.PHONY: build
build: generate
	@echo "Building ${BINARY_NAME}..."
	@mkdir -p ${BUILD_DIR}
	go build ${GOFLAGS} ${LDFLAGS} -o ${BUILD_DIR}/${BINARY_NAME} ${SOURCE_DIR}
	@echo "Build complete: ${BUILD_DIR}/${BINARY_NAME}"


# Run the application
.PHONY: run
run: build
	@echo "Running ${BINARY_NAME}..."
	${BUILD_DIR}/${BINARY_NAME} ${ARGS}

# Run with go run (without building)
.PHONY: run-direct
run-direct:
	go run ${GOFLAGS} ${LDFLAGS} ${MAIN_FILE} ${ARGS}

# Build with debug flags
.PHONY: build-debug
build-debug: GOFLAGS += -gcflags="-N -l"
build-debug: LDFLAGS += -X main.Debug=true
build-debug: build

# Build with race detector
.PHONY: build-race
build-race: GOFLAGS += -race
build-race: build

# Test
.PHONY: test
test: build
	mkdir -p ./test/tmp
	go test ./... -race
	${BUILD_DIR}/${BINARY_NAME} . -c 'g behavioural/*.mx -> foreach f { source $f }; test run *'

# Benchmark
.PHONY: bench
bench:
	go test -bench=. -benchmem ./...

# Clean build artifacts
.PHONY: clean
clean:
	@echo "Cleaning..."
	@rm -rf ${BUILD_DIR} ./test/tmp
	@rm -f coverage.out coverage.html
	@echo "Clean complete"

# Install dependencies
.PHONY: deps
deps:
	go mod download
	go mod tidy

# Lint code (requires golangci-lint)
.PHONY: lint
lint:
	golangci-lint run

# Generate code
.PHONY: generate
generate:
	@echo "Rerunning code generation..."
	go generate ./...

# Install the binary
.PHONY: install
install: build
	@echo "Installing ${BINARY_NAME}..."
	@cp ${BUILD_DIR}/${BINARY_NAME} /usr/bin/
	shell echo "/usr/bin/${BINARY_NAME}" >> /etc/shells
	@echo "Installation complete"

# Help
.PHONY: help
help:
	@echo "Available targets:"
	@echo "  make build          - Build the binary"
	@echo "  make run            - Build and run the binary"
	@echo "  make build-debug    - Build with debug symbols"
	@echo "  make test           - Run tests"
	@echo "  make bench          - Run benchmarks"
	@echo "  make clean          - Remove build artifacts"
	@echo "  make deps           - Download dependencies"
	@echo "  make lint           - Lint code"
	@echo "  make install        - Install binary to /usr/bin (requires root)"
	@echo ""
	@echo "Variables:"
#	@echo "  VERSION=x.x.x       - Set version (default: dev)"
#	@echo "  ARGS='...'          - Pass arguments to run target"
	@echo "  GOFLAGS='...'       - Additional go build flags"

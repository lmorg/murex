# Variables
BINARY_NAME=murex
GO_FLAGS=-v
BUILD_DIR=./bin
SOURCE_DIR=.

# Build variables that can be overridden
BRANCH=$(shell git rev-parse --abbrev-ref HEAD || echo "unknown")
BUILD_DATE=$(shell date -u '+%Y-%m-%d_%H:%M:%S' || echo "unknown")
LDFLAGS=-ldflags "-X github.com/lmorg/murex/app.branch=${BRANCH} -X github.com/lmorg/murex/app.buildDate=${BUILD_DATE}"
BUILD_TAGS?=$(shell cat builtins/optional/standard-opts.txt || echo "")

# Default target
.PHONY: all
all: build

# Build the binary
.PHONY: build
build: generate
	@echo "Building ${BINARY_NAME}..."
	@mkdir -p ${BUILD_DIR}
	go build ${GO_FLAGS} -tags ${BUILD_TAGS} ${LDFLAGS} -o ${BUILD_DIR}/${BINARY_NAME} ${SOURCE_DIR}
	@echo "Build complete: ${BUILD_DIR}/${BINARY_NAME}"

# Install the binary
.PHONY: install
install: build
	@echo "Installing ${BINARY_NAME}..."
	@cp ${BUILD_DIR}/${BINARY_NAME} /usr/bin/
	echo "/usr/bin/${BINARY_NAME}" >> /etc/shells
	@echo "Installation complete"

# Run the application
.PHONY: run
run: build
	@echo "Running ${BINARY_NAME}..."
	${BUILD_DIR}/${BINARY_NAME} ${ARGS}

# Build with dev flags
.PHONY: build-dev
build-dev: GO_FLAGS += -gcflags="-N -l" -race -covermode=atomic
build-dev: BUILD_TAGS += "pprof,trace,no_crash_handler"
#build-dev: LDFLAGS += -X main.Debug=true
build-debug: build

# Test
.PHONY: test
test: build
	@mkdir -p ./test/tmp
	go test ./... -count 1 -race -covermode=atomic
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

# List available build tags
.PHONY: list-build-tags
list-build-tags:
	@find . -name "*.go" -exec grep "//go:build" {} \; \
	| grep -v -E '(ignore|js|windows|linux|darwin|plan9|solaris|freebsd|openbsd|netbsd|dragonfly|aix)' \
	| sed -e 's,//go:build ,,;s,!,,;' \
	| sort -u

# Help
.PHONY: help
help:
	@echo "Available targets:"
	@echo "  make build           - Build the binary"
	@echo '  make list-build-tags - list tags supported by `$$BUILD_TAGS`'
	@echo "  make run             - Build and run the binary"
	@echo "  make build-dev       - Build with profiling and debug symbols"
	@echo "  make test            - Run tests"
	@echo "  make bench           - Run benchmarks"
	@echo "  make clean           - Remove build artifacts"
	@echo "  make deps            - Download dependencies"
	@echo "  make lint            - Lint code (requires golangci-lint)"
	@echo "  make install         - Install binary to /usr/bin (requires root)"
	@echo ""
	@echo "Variables:"
	@echo "  GO_FLAGS='...'       - Additional go build flags (default: ${GO_FLAGS})"
	@echo "  BUILD_TAGS='...'     - Additional go build tags  (default: $(shell cat ./builtins/optional/standard-opts.txt))"

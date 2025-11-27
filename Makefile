# Variables
BINARY_NAME?=murex
GO_FLAGS=-v
BUILD_DIR=./bin
SOURCE_DIR=.

# Build variables that can be overridden
BRANCH=$(shell git rev-parse --abbrev-ref HEAD || echo "unknown")
BUILD_DATE=$(shell date -u '+%Y-%m-%d_%H:%M:%S' || echo "unknown")
EXT_LDFLAGS?="-static"
LDFLAGS=-ldflags "-X github.com/lmorg/murex/app.branch=${BRANCH} -X github.com/lmorg/murex/app.buildDate=${BUILD_DATE} -extldflags=${EXT_LDFLAGS}"
BUILD_TAGS?=$(shell cat builtins/optional/standard-opts.txt || echo "")

# Default target
.PHONY: all
all: generate
all: build
all: build-wasm

# Build the binary
.PHONY: build
build:
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
run: build-dev
	@echo "Running ${BINARY_NAME}..."
	${BUILD_DIR}/${BINARY_NAME} ${ARGS}

# Build with dev flags
.PHONY: build-dev
build-dev: GO_FLAGS += -gcflags="-N -l" -race
build-dev: BUILD_TAGS := "$(BUILD_TAGS),pprof,trace,no_crash_handler"
build-dev: build

# Build with TinyGo instead of default Go binary
.PHONY: build-tinygo
build-tinygo: LDFLAGS=-ldflags "-X github.com/lmorg/murex/app.branch=${BRANCH} -X github.com/lmorg/murex/app.buildDate=${BUILD_DATE}"
build-tinygo: BUILD_TAGS := "$(BUILD_TAGS),tinygo,no_cachedb,no_cmd_select,no_pty"
build-tinygo:
	@echo "Building ${BINARY_NAME} using TinyGo..."
	tinygo build -x -tags ${BUILD_TAGS}

# Build WASM
.PHONY: build-wasm
build-wasm: export GOOS=js
build-wasm: export GOARCH=wasm
build-wasm: build-tinygo

# Test
.PHONY: test
test:
	@mkdir -p ./test/tmp
	go test ./... -count 1 -race -covermode=atomic -tags "$(BUILD_TAGS),no_crash_handler"
	${BUILD_DIR}/${BINARY_NAME} -c 'g behavioural/*.mx -> foreach f { source $$f }; test run *'

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

# Update dependencies
.PHONY: update-deps
update-deps:
	go get -u ./...
	go mod tidy

# Lint code (requires golangci-lint)
.PHONY: lint
lint:
	golangci-lint run

.PHONY: build-docgen
build-docgen:
	@echo "Building docgen..."
	go build -v ./utils/docgen

# Generate code
.PHONY: generate
generate:
	@echo "Rerunning code generation..."
	go generate ./...

# Generate animated gifs
.PHONY: update-vhs
update-vhs: build
	@echo "Rerunning vhs image generation..."
	docgen -verbose -warning -config gen/vhs/docgen.yaml
	@$(foreach f,$(wildcard gen/vhs/generated/*.tape),vhs ${f};)
	git add images/vhs* gen/vhs gen/vuepress/styles/images.scss
	git commit -m "vhs: content updated" --no-verify

# List available build tags
.PHONY: list-build-tags
list-build-tags:
	@find . -name "*.go" -exec grep "//go:build" {} \; \
	| grep -v -E '(ignore|js|windows|linux|darwin|plan9|solaris|freebsd|openbsd|netbsd|dragonfly|aix)' \
	| sed -e 's,//go:build ,,;s,!,,;' \
	| sort -u
	@echo "sqlite_omit_load_extension\nosusergo\nnetgo"

# readline package development
local_readline = local/readline

.PHONY: 
local-dev-readline:
ifneq "$(wildcard $(local_readline)/.)" ""
	cd $(local_readline)
	git pull
else
	@mkdir -p local
	git clone git@github.com:lmorg/readline.git $(local_readline)
endif
	cd $(local_readline)
	go mod edit -replace "github.com/lmorg/readline/v4=./$(local_readline)"
	go mod tidy
	@echo ""
	@echo "Before you push any changes of Murex, you will need to run:"
	@echo "    make remote-readline"

.PHONY:
remote-readline:
	go mod edit -dropreplace=github.com/lmorg/readline/v4
	go mod tidy

# Help
.PHONY: help
help:
	@echo "Available commands:"
	@echo "  make build           - Build the Murex"
	@echo "  make install         - Install Murex to /usr/bin (requires root)"
	@echo '  make list-build-tags - list tags supported by `$$BUILD_TAGS`'
	@echo "  make clean           - Remove build artifacts"
	@echo ""
	@echo "Development tools:"
	@echo "  make build-dev       - Build with profiling and debug symbols"
	@echo "  make run             - Build and run a dev build of Murex"
	@echo "  make test            - Run tests"
	@echo "  make bench           - Run benchmarks"
	@echo "  make update-deps     - Update all Go dependencies"
	@echo "  make update-vhs      - Update all VHS-generated animated gifs"
	@echo "  make lint            - Lint code (requires golangci-lint)"
	@echo ""
	@echo "Variables:"
	@echo "  GO_FLAGS='...'       - Additional go build flags (default: ${GO_FLAGS})"
	@echo "  BUILD_TAGS='...'     - Additional go build tags  (default: $(shell cat ./builtins/optional/standard-opts.txt))"

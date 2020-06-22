#!/bin/bash

set -ev

test/pre-commit --no-cd

echo "Starting count server...."
export MUREX_TEST_COUNT=http
go run github.com/lmorg/murex/test/count/server 2>/dev/null &
sleep 1
        
echo "Run golang unit tests...."
go test ./... -count 1 -race -coverprofile=coverage.txt -covermode=atomic
export MUREXTESTS="$(curl -s http://localhost:38000/t)"
echo "$MUREXTESTS tests completed"

echo "Compiling murex...."
go install github.com/lmorg/murex

echo "Run murex shell script unit tests...."
murex --run-tests

echo "Building latest binaries...."
murex ./test/build_all_platforms.mx $MUREX_BUILD_FLAGS

echo "Fin!"
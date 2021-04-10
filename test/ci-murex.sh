#!/bin/bash

set -ev

#test/pre-commit --no-cd

. /etc/ci-murex.env

echo "Compiling stringer...."
go build -o /bin/stringer golang.org/x/tools/cmd/stringer 

echo "Updating auto-generated code...."
go generate ./...

echo "Compiling docgen...."
go install github.com/lmorg/murex/utils/docgen

echo "Compiling murex docs...."
docgen -config gen/docgen.yaml

echo "Compiling murex...."
go install github.com/lmorg/murex

echo "Starting count server...."
export MUREX_TEST_COUNT=http
go run github.com/lmorg/murex/test/count/server 2>/dev/null &
sleep 1
        
echo "Run golang unit tests...."
go test ./... -count 1 -race -coverprofile=coverage.txt -covermode=atomic
curl -s http://localhost:38000/t > ./murex-test-count.txt
echo "$(cat ./murex-test-count.txt) tests completed"

echo "Run murex shell script unit tests...."
murex --run-tests

echo "Run murex flag unit tests...."
./murex -c 'source: ./flags_test.mx; try {test: run *}'

echo "Fin!"
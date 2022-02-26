#!/bin/bash

set -ev

. /etc/ci-murex.env

#echo "Compiling stringer...."
#go build -o /bin/stringer golang.org/x/tools/cmd/stringer 
#
#echo "Updating auto-generated code...."
#go generate ./...

echo "Compiling docgen...."
go install github.com/lmorg/murex/utils/docgen

echo "Compiling murex docs...."
docgen -config gen/docgen.yaml

echo "Compiling murex...."
go install github.com/lmorg/murex

echo "Starting count server...."
export MUREX_TEST_COUNT=http
go run github.com/lmorg/murex/test/count/server 2>/dev/null &
sleep 3

echo "Running golang unit tests...."
mkdir -p ./test/tmp
go test ./... -count 1 -race -coverprofile=coverage.txt -covermode=atomic
curl -s http://localhost:38000/t > ./murex-test-count.txt
echo "$(cat ./murex-test-count.txt) tests completed"

echo "Running murex behavioural tests...."
murex -c 'g: behavioural/* -> foreach: f { source $f }; try {test: run *}'

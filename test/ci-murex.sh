#!/bin/bash

set -ev

. /etc/ci-murex.env

export MUREXVERSION="$(murex -c 'version --no-app-name')"
OLDVER="$(curl -s https://murex.rocks/VERSION | head -n1)"

if [ "$MUREXVERSION" == "$OLDVER" ]; then
    echo "No version change, skipping tests."
else
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

    echo "Running murex behavioural tests...."
    murex -c 'g: behavioural/* -> foreach: f { source $f }; try {test: run *}'
fi
#!/bin/bash

set -ev

. /etc/ci-murex.env

echo "Compiling docgen...."
go install github.com/lmorg/murex/utils/docgen

echo "Compiling murex docs...."
docgen -config gen/docgen.yaml

echo "Compiling murex...."
go install github.com/lmorg/murex


export MUREXVERSION="$(murex -c 'version --no-app-name')"
OLDVER="$(curl -s https://murex.rocks/VERSION | head -n1)"

if [ "$MUREXVERSION" == "$OLDVER" ]; then
    echo "No version change, skipping tests."
else
    echo "Running murex behavioural tests...."
    murex -c 'g: behavioural/*.mx -> foreach: f { source $f }; test: run *'
fi

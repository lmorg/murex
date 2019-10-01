#!/bin/bash

set -ev

test/pre-commit

echo "Starting count server...."
export MUREX_TEST_COUNT=http
go run github.com/lmorg/murex/test/count/server 2>/dev/null &
sleep 1
        
echo "Run golang unit tests...."
go test ./... -count 10 -race -coverprofile=coverage.txt -covermode=atomic
export MUREXTESTS="$(curl -s http://localhost:38000/t)"

echo "Compiling murex...."
go install github.com/lmorg/murex

echo "Run murex shell script unit tests...."
murex --run-tests

echo "Building latest binaries...."
murex ./test/build_all_platforms.mx --no-colour --inc-latest --compress

echo "Building website...."
export MUREXVERSION="$(murex -c 'version --no-app-name')"
export MUREXCOMMITS="$(git rev-parse HEAD)"
export MUREXCOMMITL="$(git rev-parse HEAD)"

sed -i "s/\$DATE/`date`/;
        s/\$COMMITHASHSHORT/$MUREXCOMMITS/;
        s/\$COMMITHASHLONG/$MUREXCOMMITL/;
        s/\$MUREXVERSION/$MUREXVERSION/;
        s/\$MUREXTESTS/$MUREXTESTS/" \
        gen/website/footer.html

for f in *.md; do
        gen/website/find-exec.sh $f
done
find docs -name "*.md" -exec gen/website/find-exec.sh {} \;

mkdir /website
mv -v *.html gen/website/*.css ./bin ./docs /website/

#!/bin/sh

test/pre-commit

echo "Starting count server...."
go run github.com/lmorg/murex/test/count/server &
export MUREX_TEST_COUNT=http

echo "Run golang unit tests...."
go test ./... -count 10 -race -coverprofile=coverage.txt -covermode=atomic
export MUREXTESTS=$(curl http://localhost:38000/t)
export MUREXVERSION="$(./murex -c 'version --no-app-name')"

echo "Run murex shell script unit tests..."
murex --run-tests

echo "Building latest binaries...."
./murex ./test/build_all_platforms.mx --no-colour --inc-latest --compress

echo "Building website...."
mv -v ./bin ./docs/
sed -i "s/\$DATE/`date`/; s/\$COMMITHASHSHORT/`git rev-parse HEAD | cut -c1-7`/; s/\$COMMITHASHLONG/`git rev-parse HEAD`/; s/\$MUREXVERSION/$MUREXVERSION/; s/\$MUREXTESTS/$MUREXTESTS/" gen/website/footer.html
mv -v *.md ./docs/
mv -v gen/website/*.css ./docs/
find docs -name "*.md" -exec gen/website/find-exec.sh {} \;

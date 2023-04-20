#!/bin/bash

set -ev

. /etc/ci-murex.env

mkdir -p /website


export MUREXVERSION="$(murex -c 'version --no-app-name')"
#export MUREXVERSION="$(cat app/app.go | grep 'const Version' | grep -E -o '[0-9]+\.[0-9]+\.[0-9]+')"
OLDVER="$(curl -s https://murex.rocks/VERSION | head -n1)"

if [ "$MUREXVERSION" == "$OLDVER" ]; then
        echo "No version change, skipping binary build."
else
        echo "Building latest binaries...."
        murex ./test/build_all_platforms.mx $MUREX_BUILD_FLAGS
        mv -v ./bin /website/

        echo "Compiling WebAssembly...."
        export GOOS=js
        export GOARCH=wasm
        go build -o ./gen/website/wasm/murex.wasm
        cp -v "$(go env GOROOT)/misc/wasm/wasm_exec.js" ./gen/website/wasm/
fi

echo "Building website...."

export MUREXCOMMITS="$(git rev-parse HEAD | cut -c1-7)"
export MUREXCOMMITL="$(git rev-parse HEAD)"
export MUREXTESTS="$(cat ./test/murex-test-count.txt)"

sed -i "s/\$DATE/`date`/g;
        s/\$COMMITHASHSHORT/$MUREXCOMMITS/g;
        s/\$COMMITHASHLONG/$MUREXCOMMITL/g;
        s/\$MUREXVERSION/$MUREXVERSION/g;
        s/\$MUREXTESTS/$MUREXTESTS/g" \
        gen/website/header.html

sed -i "s/\$DATE/`date`/;
        s/\$COMMITHASHSHORT/$MUREXCOMMITS/g;
        s/\$COMMITHASHLONG/$MUREXCOMMITL/g;
        s/\$MUREXVERSION/$MUREXVERSION/g;
        s/\$MUREXTESTS/$MUREXTESTS/g" \
        gen/website/footer.html

cp gen/website/404.md .
for f in *.md; do
        murex gen/website/find-exec.mx $f
done
find docs -name "*.md" -exec gen/website/murex gen/website/find-exec.mx {} \;

sed -i '0,/<img src/s//<img class="no-border" src/;
        0,/<img src/s//<img class="no-border" src/;
        0,/<img src/s//<img class="no-border" src/;' \
        README.html

sed -i '0,/<img src/s//<img class="no-border" src/;
        0,/<img src/s//<img class="no-border" src/;
        0,/<img src/s//<img class="no-border" src/;
        0,/<img src/s//<img class="no-border" src/;' \
        INSTALL.html

sed -i '0,/<img src/s//<img class="no-border" src/;' \
        DOWNLOAD.html

sed -i 's.\\|.|.g;' \
        docs/user-guide/rosetta-stone.html

echo "$MUREXVERSION" > VERSION

mv VERSION *.html *.svg gen/website/assets/* ./docs ./images /website/

echo "Fin!"

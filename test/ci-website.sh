#!/bin/bash

set -ev

. /etc/ci-murex.env

#[ -z "$MUREX_BUILD_WEBSITE_ONLY" ] && \
#        export MUREX_BUILD_WEBSITE_ONLY="$(git describe --tags --match website)"

#export MUREX_BUILD="$(git log -1 --pretty=%B | egrep '^website: ')"

mkdir -p /website

export MUREXVERSION="$(murex -c 'version --no-app-name')"
OLDVER="$(curl -s https://murex.rocks/VERSION | head -n1)"

if [ "$MUREXVERSION" == "$OLDVER" ]; then
        echo "No version change, skipping binary build."
else
        echo "Building latest binaries...."
        murex ./test/build_all_platforms.mx $MUREX_BUILD_FLAGS
        mv -v ./bin /website/
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
        gen/website/find-exec.sh $f
done
find docs -name "*.md" -exec gen/website/find-exec.sh {} \;

sed -i '0,/<img src/s//<img class="no-border" src/;
        0,/<img src/s//<img class="no-border" src/;
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

echo "Compiling WebAssembly...."

export GOOS=js
export GOARCH=wasm
go build -o ./gen/website/wasm/murex.wasm
cp -v "$(go env GOROOT)/misc/wasm/wasm_exec.js" ./gen/website/wasm/

echo "$MUREXVERSION" > VERSION

mv VERSION *.html *.svg gen/website/assets/* ./docs /website/ 

echo "Fin!"
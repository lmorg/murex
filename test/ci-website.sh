#!/bin/bash

set -ev

. /etc/ci-murex.env

#[ -z "$MUREX_BUILD_WEBSITE_ONLY" ] && \
#        export MUREX_BUILD_WEBSITE_ONLY="$(git describe --tags --match website)"

#export MUREX_BUILD="$(git log -1 --pretty=%B | egrep '^website: ')"

mkdir -p /website


echo "Building latest binaries...."
#if [ "$MUREX_BUILD" != "website: " ]; then
        murex ./test/build_all_platforms.mx $MUREX_BUILD_FLAGS
        mv -v ./bin /website/
#else
#        echo "!!! Commit appears to only update website, skipping compile tests !!!"
#fi

echo "Building website...."
export MUREXVERSION="$(murex -c 'version --no-app-name')"
export MUREXCOMMITS="$(git rev-parse HEAD | cut -c1-7)"
export MUREXCOMMITL="$(git rev-parse HEAD)"
export MUREXTESTS="$(cat ./murex-test-count.txt)"

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
        0,/<img src/s//<img class="no-border" src/;' \
        README.html

echo "Compiling WebAssembly...."

export GOOS=js
export GOARCH=wasm
go build -o ./gen/website/wasm/murex.wasm
cp -v "$(go env GOROOT)/misc/wasm/wasm_exec.js" ./gen/website/wasm/



mv *.html gen/website/assets/* ./docs /website/

echo "Fin!"
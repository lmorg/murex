#!/bin/bash

#[ -z "$MUREX_BUILD_WEBSITE_ONLY" ] && \
#        export MUREX_BUILD_WEBSITE_ONLY="$(git describe --tags --match website)"

export MUREX_BUILD="$(git log -1 --pretty=%B | egrep '^website: ')"

mkdir -p /website


echo "Building latest binaries...."
if [ "$MUREX_BUILD" != "website: " ]; then
        murex ./test/build_all_platforms.mx $MUREX_BUILD_FLAGS
        mv ./bin /website/
else
        echo "!!! Commit appears to only update website, skipping compile tests !!!"
fi

echo "Building website...."
export MUREXVERSION="$(murex -c 'version --no-app-name')"
export MUREXCOMMITS="$(git rev-parse HEAD | cut -c1-7)"
export MUREXCOMMITL="$(git rev-parse HEAD)"
export MUREXTESTS="$(cat ./murex-test-count.txt)"

sed -i "s/\$DATE/`date`/;
        s/\$COMMITHASHSHORT/$MUREXCOMMITS/;
        s/\$COMMITHASHLONG/$MUREXCOMMITL/;
        s/\$MUREXVERSION/$MUREXVERSION/;
        s/\$MUREXTESTS/$MUREXTESTS/" \
        gen/website/footer.html

cp gen/website/404.md .
for f in *.md; do
        gen/website/find-exec.sh $f
done
find docs -name "*.md" -exec gen/website/find-exec.sh {} \;



echo "Compiling WASM...."

export GOOS=js
export GOARCH=wasm
go build -o ./gen/website/wasm/murex.wasm
cp -v "$(go env GOROOT)/misc/wasm/wasm_exec.js" ./gen/website/wasm/

mv *.html gen/website/assets/* ./docs /website/



echo "Fin!"
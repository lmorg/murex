#!/bin/bash

[ -z "$MUREX_BUILD_BIN" ]     && \
        export MUREX_BUILD_BIN="$(git describe --tags --match bin)"
[ -z "$MUREX_BUILD_WEBSITE" ] && \
        export MUREX_BUILD_WEBSITE="$(git describe --tags --match website)"

mkdir -p /website


echo "Building latest binaries...."
#if [ "$MUREX_BUILD_BIN" == "bin" ]; then
        murex ./test/build_all_platforms.mx $MUREX_BUILD_FLAGS
        mv ./bin /website/
#fi

echo "Building website...."
#if [ "$MUREX_BUILD_WEBSITE" == "website" ]; then
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

        mv *.html gen/website/assets/* ./docs /website/
#fi

echo "Fin!"
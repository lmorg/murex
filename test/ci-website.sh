#!/bin/bash

echo "Building latest binaries...."
murex ./test/build_all_platforms.mx $MUREX_BUILD_FLAGS

echo "Building website...."
export MUREXVERSION="$(murex -c 'version --no-app-name')"
export MUREXCOMMITS="$(git rev-parse HEAD | cut -c1-7)"
export MUREXCOMMITL="$(git rev-parse HEAD)"

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

mkdir /website | true
mv -v *.html gen/website/*.css ./bin ./docs /website/
echo "Fin!"
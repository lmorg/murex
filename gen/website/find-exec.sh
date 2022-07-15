#!/bin/sh

#export MUREXVERSION="$(cat app/app.go | grep 'const Version' | egrep -o '[0-9]+\.[0-9]+\.[0-9]+')"

pandoc "$1" > "$1.tmp"

# html file extension
html=$(printf "$1" | sed 's/\.md/.html/')

# replace all hyperlinks to .md with .html
sed -i 's/\.md/.html/g;
        s/<li><p>/<li>/;
        s,</p></li>,</li>,;
        s/version\.svg/version.svg\?v='"$MUREXVERSION/" "$1.tmp"

cat gen/website/header.html "$1.tmp" gen/website/footer.html > "$html"

rm "$1.tmp" "$1"
#!/bin/sh

pandoc "$1" > "$1.tmp"

# html file extension
html=$(printf "$1" | sed 's/\.md/.html/')

# replace all hyperlinks to .md with .html
sed -i 's/\.md/.html/' "$1.tmp"

cat test/docs/header.html "$1.tmp" test/docs/footer.html > "$html"

rm "$1.tmp" 
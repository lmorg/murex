#!/bin/sh

pandoc "$1" > "$1.tmp"

# html file extension
html=$(printf "$1" | sed 's/\.md/.html/')

# replace all hyperlinks to .md with .html
sed -i 's/\.md/.html/; s/<li><p>/<li>/; s,</p></li>,</li>,' "$1.tmp"

cat gen/website/header.html "$1.tmp" gen/website/footer.html > "$html"

rm "$1.tmp" "$1"
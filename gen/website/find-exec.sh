#!/bin/env murex

pandoc $1 |> $1.tmp

# html file extension
$1 -> regexp 's/\.md/.html/' -> set html

# replace all hyperlinks to .md with .html
sed -i %(s/\.md/.html/g;
         s/<li><p>/<li>/;
         s,</p></li>,</li>,;
         s/version\.svg/version.svg\?v=$MUREXVERSION/) $1.tmp

cat gen/website/header.html $1.tmp gen/website/footer.html |> $html

open $1 -> regexp %(f,<h1>(.*?)</h1>,) -> set title

sed -i "s,<title></title>,<title>$title - Murex Shell</title>," $html

rm $1.tmp $1
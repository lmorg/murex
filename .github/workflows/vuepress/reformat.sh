#!/bin/bash

filename="$1"

sed -r -i 's,<h2>Table of Contents</h2>,,' $filename

lower="$(echo \"$filename\" | tr '[:upper:]' '[:lower:]')"
if [ "$filename" = "$lower" ]; then
    mv "$filename" "$lower"
    filename="$lower"
fi 
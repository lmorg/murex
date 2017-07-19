#!/bin/sh

MANPATH=$(man -w | sed -r 's/:/\n/g')
WRITEPATH=$(pwd)

rm *.txt

for d in $MANPATH; do
    cd $d
    #for d2 in $(ls | grep man); do
    for d2 in man{1,6,7,8}; do
        cd $d2 2>/dev/null
        for file in $(ls *.gz); do
            echo $d/$d2/$file...
            cmd=$(echo $file | sed -r 's/\..*//')
            zcat $file | egrep -o '\\fB.*\\fR' | sed -r 's#(\\fB|\\fR)##g; s/, /\n/g; s/\\//g; s/[=\[].*$//g' | sort | uniq >> $WRITEPATH/$cmd.txt
        done
        cd ..
    done
done

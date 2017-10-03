#!/bin/sh
P=$(pwd)
OUT=$P/../therealplato.github.io
TMP=$P/tmpgit
mv $OUT/.git $TMP
mv $OUT/CNAME .
rm -rf $OUT/*
hugo -s hugosite -d $OUT
touch $OUT/.nojekyll
mv $TMP $OUT/.git
mv CNAME $OUT/CNAME

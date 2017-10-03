#!/bin/sh
P=$(pwd)
OUT=$P/../therealplato.github.io
TMP=$P/tmpgit
mv $OUT/.git $TMP
rm -rf $OUT/*
hugo -s hugosite -d $OUT
touch $OUT/.nojekyll
mv $TMP $OUT/.git

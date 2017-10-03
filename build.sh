#!/bin/sh
P=$(pwd)
OUT=$P/../therealplato.github.io
rm -rf $OUT/*
hugo -s hugosite -d $OUT

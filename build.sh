#!/bin/sh
P=$(pwd)
rm -rf docs
hugo -s hugosite -d $P/docs

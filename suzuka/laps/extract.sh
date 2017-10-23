#!/bin/bash
years=(1996)
# years=(1996 1997 1998 1999 2000 2001 2002 2003 2004 2005 2006 2007 2008 2009 2010 2011 2012 2013 2014 2015 2016 2017)

for y in ${years[@]}; do
  if [ -f $y.1.json ]; then
    cat <(cat $y.json| jq '.MRData .RaceTable .Races[] .Laps[]') <(cat $y.1.json| jq '.MRData .RaceTable .Races[] .Laps[]') | jq -s '.' > $y.laps.json
  else
    cat $y.json | jq '.MRData .RaceTable .Races[] .Laps[]' | jq -s '.' > $y.laps.json
  fi
done

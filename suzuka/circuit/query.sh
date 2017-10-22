#!/bin/bash
# look up Suzuka race index by year:
years=(1996 1997 1998 1999 2000 2001 2002 2003 2004 2005 2006 2007 2008 2009 2010 2011 2012 2013 2014 2015 2016 2017)
for y in ${years[@]}; do
  # respect API rate limit:
  if [ ! -f $y.json ]; then
    curl -so $y.json http://ergast.com/api/f1/$y/circuits.json
    sleep 1
  fi
  i=$(cat $y.json | jq '.MRData .CircuitTable .Circuits |map(.circuitId) | index("suzuka")')
  echo "http://ergast.com/api/f1/$y/$i/laps.json"
done

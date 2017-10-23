#!/bin/bash
urls=(
http://ergast.com/api/f1/1999/14/laps
http://ergast.com/api/f1/2000/15/laps
http://ergast.com/api/f1/2004/16/laps
http://ergast.com/api/f1/2005/17/laps
http://ergast.com/api/f1/2006/16/laps
http://ergast.com/api/f1/2009/14/laps
http://ergast.com/api/f1/2010/14/laps
http://ergast.com/api/f1/2011/14/laps
http://ergast.com/api/f1/2012/15/laps
http://ergast.com/api/f1/2013/15/laps
http://ergast.com/api/f1/2014/16/laps
http://ergast.com/api/f1/2016/18/laps
)
testurls=(
http://ergast.com/api/f1/1999/14/laps
)

for u in ${urls[@]}; do
	echo $u
	y=$(echo $u | grep -oE "[0-9]{4}")
  if [ ! -f $y.1.json ]; then
    curl -so $y.1.json "$u.json?limit=500&offset=1000"
    #respect API rate limit:
		sleep 1
  fi
done

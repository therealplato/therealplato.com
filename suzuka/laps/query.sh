#!/bin/bash
urls=(
http://ergast.com/api/f1/1996/14/laps
http://ergast.com/api/f1/1997/15/laps
http://ergast.com/api/f1/1998/14/laps
http://ergast.com/api/f1/1999/14/laps
http://ergast.com/api/f1/2000/15/laps
http://ergast.com/api/f1/2001/15/laps
http://ergast.com/api/f1/2002/15/laps
http://ergast.com/api/f1/2003/14/laps
http://ergast.com/api/f1/2004/16/laps
http://ergast.com/api/f1/2005/17/laps
http://ergast.com/api/f1/2006/16/laps
http://ergast.com/api/f1/2009/14/laps
http://ergast.com/api/f1/2010/14/laps
http://ergast.com/api/f1/2011/14/laps
http://ergast.com/api/f1/2012/15/laps
http://ergast.com/api/f1/2013/15/laps
http://ergast.com/api/f1/2014/16/laps
http://ergast.com/api/f1/2015/16/laps
http://ergast.com/api/f1/2016/18/laps
http://ergast.com/api/f1/2017/17/laps
)
testurls=(
http://ergast.com/api/f1/1996/14/laps
)

for u in ${urls[@]}; do
	echo $u
	y=$(echo $u | grep -oE "[0-9]{4}")
  #respect API rate limit:
  if [ ! -f $y.json ]; then
    curl -o $y.json $u.json?limit=1000
		sleep 1
  fi
done

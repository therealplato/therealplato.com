---
title: "#soberoctober day 22: preparing to chart F1 data"
date: 2017-10-22T22:25:25+13:00
draft: false
---

I'd like to know the answer to "How did lap times evolve throughout the years of the F1 Suzuka Grand Prix?"

I found an excellent data source: the
[ergast.com API](http://ergast.com/mrd/methods/laps/). The "lap times" docs
indicate that this dataset is available for races since 1996. First I'll look
up the index of the Suzuka GP for each year, since sometimes circuits get added
or dropped:

```
#!/bin/bash
# find Suzuka race number:
years=(1996 1997 1998 1999 2000 2001 2002 2003 2004 2005 2006 2007 2008 2009 2010 2011 2012 2013 2014 2015 2016 2017)
for y in ${years[@]}; do
  # respect the API by not redownloading and rate limiting:
  if [ ! -f $y.json ]; then
    curl -so $y.json http://ergast.com/api/f1/$y/circuits.json
    sleep 1
  fi
  i=$(cat $y.json | jq '.MRData .CircuitTable .Circuits |map(.circuitId) | index("suzuka")')
  echo "http://ergast.com/api/f1/$y/$i/laps"
done
```
From the output, I note that that 2007 and 2008 have an index of `null`. I
wonder why the races were not run? Maybe that was the "Fuji" circuit I've heard
of? All other years, it's either the 14th, 15th or 16th race of the season.

```
# get lap times for the race
curl -o 1996.json http://ergast.com/api/f1/1996/14/laps?limit=1000
```

The 1996 race has a lot of names I recognize! Shumacher, Brundle, Herbert, Verstappen... Also, it looks like very few drivers finished.

Several races have a lap count exceeding the API query limit of 1000 so I need to make additional queries and join the datasets. I'll manually copy and paste the second page of results into the first page; I don't think it's worth automating.

If I do this again I'll run local sql queries against [this docker image](https://github.com/jcnewell/ergast-f1-api) instead of wrangling JSON.

*edit* This ended up taking hours of struggling with JQ. I stubbornly refused to try docker after seeing they don't release linux versions
and Arch's requires screwing with kernel modules. The hardest bit was figuring out how to make JQ turn a sequence of objects into an array
of those objects. It's the "slurp" flag.
```
cat <(cat 1999.json| jq '.MRData .RaceTable .Races[] .Laps[]') <(cat 1999.1.json| jq '.MRData .RaceTable .Races[] .Laps[]') | jq -s '.' > 1999.laps.json
```

### Bonus Arch Linux hack:
Since Arch comes with basically nothing, you have to install everything. And sometimes it's not obvious where it lives. I wrote an alias
that searches Pacman's database of not-yet-installed packages for the command you're trying to find:

```
 suzuka/circuit Ω pacwut jq
usr/bin/jq is owned by community/jq 1.5-5

 suzuka/circuit Ω which pacwut
pacwut () {
        pacman -Fo $@
        pacman -Fo /bin/$@
        pacman -Fo /usr/bin/$@
        pacman -Fo /usr/local/bin/$@
}
 suzuka/circuit Ω sudo pacman -S community/jq
```

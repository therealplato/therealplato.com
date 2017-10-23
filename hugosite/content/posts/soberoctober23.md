---
title: "#soberoctober day 23: charting f1"
date: 2017-10-23T12:19:02+13:00
draft: false
---

Continuing from yesterday, I've got a new hypothesis. If I plot the lap times of one driver throughout a race, I think old races (when
pitstops included 
refueling) will show a gentle-drop-steep-climb sawtooth pattern, and new races will show a slow-climb-fast-drop pattern with lower
amplitude. I expect this because cars are slower when they have to carry fuel, but they also get slower as tires wear out. I think for the
old races, the fuel weight will outweigh the fresh tires.

D3 pseudocode:
```
Chart 1, fastest lap history
## Data:
fastest = <race, driver, lap number, lap time>
for r in races
  for l in r.laps
    for d in l.drivers
      t = d.laptime
      if t < fastest.t; fastest = <r,d,l,t>

## Chart:
line chart, year vs fastest
hover a data point to see driver, lap number, lap time

enter:
create svg circle 
update:
place svg circle at tx, ty
leave: 
destroy svg circle
```

<script src="https://d3js.org/d3.v4.js"></script>
<script>
  // All D3 code goes here. For Example:
  console.log(d3.version); // 4.7.4
</script>

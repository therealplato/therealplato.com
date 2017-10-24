---
title: "#soberoctober day 24"
date: 2017-10-24T21:41:48+13:00
draft: false
---
<script src="https://d3js.org/d3.v4.js"></script>
<script src="/js/soberoctober24.js"></script>

I was gonna make a sweet-ass graph that put yesterday's to shame but then I watched some youtube and now it's 21:45. I'll give it a go anyway!

```
# get timing data for one driver's race:
suzuka/laps Ω cat 2001.laps.json |\
jq '.[] .Timings[] |select(.driverId == "alonso") .time' |\
jq -s '.' > ../../hugosite/static/misc/2001.alonso.laps.json
```

<div id="laptimes-container"></div>

_@23:27_ - Meh. I'll take it. I learned how to deal with margins, how to draw and position axes, how to zero pad numbers. 

[chart source](https://github.com/therealplato/therealplato.com/blob/e171f05a3b7a05351e4c33a474290bd12df0b5fa/hugosite/static/js/soberoctober24.js)

Next time: plotting a second time series on the same chart; show details of the data points on hover; make x-axis start at 1.

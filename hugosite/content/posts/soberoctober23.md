---
title: "#soberoctober day 23: charting f1 (barely)"
date: 2017-10-23T12:19:02+13:00
draft: false
---
Continuing from [yesterday](/posts/soberoctober22), I've got a new hypothesis. The lap times of individual drivers will follow a sawtooth pattern throughout the race, as tires wear out and pitstops replace them with faster fresh tires. Additionally, refueling was allowed in past years. I think that the deviation between slow and fast laps will be lower in the refueling races than in post-refueling races. (Because the modern cars have to carry the whole race's worth of fuel from the start, they will show a more significant difference between early slow laps and late fast laps, than the refueling cars.)
<hr>
OK scratch that lol, I hoped to gain some D3 skills but instead I learned how to wrangle data. It might be a good idea to lower my expectations.

I spent many more hours of trial and error to combine multiple pages from the Ergast API, then extract individual lap timing data objects, and then an unholy amount of googling until I found jq's `--slurp` option to input objects and output an array wrapping them. The result was [three shell scripts and one go script](https://github.com/therealplato/therealplato.com/tree/master/suzuka/laps).

Once I had data, learning elementary [d3 code](https://github.com/therealplato/therealplato.com/blob/bcc0bc3e4d692a75c26eacdc4a9f15752487818e/hugosite/static/js/soberoctober23.js) was a breeze, clocking in at only 40 minutes or so. Yes, I know it's [ugly](/images/dwi.gif).

Historic Formula 1 Suzuka Grand Prix lap time records:
<div id="laptimes-container"></div>
<script src="https://d3js.org/d3.v4.js"></script>
<script src="/js/soberoctober23.js"></script>


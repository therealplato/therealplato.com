---
title: "Soberoctober05"
date: 2017-10-04T21:23:04+13:00
draft: true
---

 
While trying to use two finger scrolling, I observed this entertaining behavior:
```gherkin
Given my cursor is on a link
When I touch one finger
And then a second finger
And then drag up or down
Then each scroll event opens a new tab
```
Thank God for "Close tabs to the right" as this behavior can easily open 50 copies of the link under the cursor.


```sh
 ~ Ω batt  
0.56

 ~ Ω which batt
batt () {
        X=$(cat /sys/class/power_supply/BAT0/energy_now) 
        Y=$(cat /sys/class/power_supply/BAT0/energy_full) 
        python2 -c "z=$X/$Y.; print('%.2f' % z)"
}
```

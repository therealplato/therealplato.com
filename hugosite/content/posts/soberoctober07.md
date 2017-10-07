---
title: "#soberoctober day 7: one of us"
date: 2017-10-07T18:26:49+13:00
draft: false
---

Today's societies are impersonal behemoths. Our governments attempt to apply identical rules to millions or billions of people with very
different lives, and our media and education systems help by reinforcing expected patterns of behavior.

I've deleted paragraphs from this post out of fear that society at large would punish me for sharing those bits of information. I'm not sure
if those bits would result in legal consequences; if they would trigger individuals to dox, swat or DDOS me; or if no one would
care. There are rules - both laws and social conventions - that are enforced, even though I'm not aware of them or haven't agreed to them.

That's unjust, and gets worse when those rules are applied differently to privileged individuals.

My dream society is *a tribe of individuals who have explicitly agreed to a set of clearly defined rules and consequences for breaking
them*.

We have sufficient technology now to operate unlimited, low-cost digital tribes, with any combination of rules that humans desire to live by.

For the rest of #soberoctober, I intend to explore this space.

### how to check battery level on arch linux
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

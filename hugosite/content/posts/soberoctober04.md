---
title: "#soberoctober day 4: a bit of serverless"
date: 2017-10-04T19:34:11+13:00
draft: true
---

## Today's project: try out Amazon Lambda
I first heard of "serverless" computer at Full Stack Fest 2016. [This story](https://www.youtube.com/watch?v=9IrFIobZUEA&t=22m10s) really
drove home the point that it's order of magnitudes cheaper for many use cases (especially event-driven "jobs" that don't need continuous CPU
use.) 

## Just Arch Things
Haven't found any obvious way to make my volume keys work. Apparently I can manually find the keycodes and edit xorg configuration
files to map them to XF86 audio events. F dat, `alsamixer` command does the trick for now.

While trying to use two finger scrolling, I observed this entertaining behavior:
```gherkin
Given my cursor is on a link
When I touch one finger
And then a second finger
And then drag up or down
Then each scroll event opens a new tab
```
Thank God for "Close tabs to the right" as this behavior can easily open 50 copies of the link under the cursor.

*thatsthejoke.jpg*: `cool-retro-term` abbreviates to `CRT` 

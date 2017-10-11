---
title: "#soberoctober day 11: CLI timezone conversion"
date: 2017-10-11T22:12:35+13:00
draft: false
---

Movio has offices in US, UK, DE, NZ and has global clients. I frequently find myself asking "Is it office hours in the US? Is it the middle
of the night in Asia?" We keep records with UTC timestamps (as all programmers should!) but it's still irritating to to conceptualize and
convert more than two timezones. 

In the past, I've used several helpful tools:

* The excellent [timeanddate.com converter](https://www.timeanddate.com/worldclock/converter.html?iso=20171010T100100&p1=22&p2=tz_pt) gives you multiple zones and the
ability to type or drag to set any zone to any time.
* I've used my Android 6.x phone as a reference, because its stock clock app supports multiple timezones.
* OSX does not support multiple clocks, wtf? I have recently installed the third party "Hour - World Clock" from App Store, which gives me a multi-clock dropdown from my OSX taskbar.

I like the portability and scriptability of CLI solutions so I have started building a tool to answer my timezone questions.

Introducing [`timez`](https://github.com/therealplato/timez):
```
~ Ω go get -u github.com/therealplato/timez
~ Ω timez
Pacific/Auckland: 2017-10-11 22:35:39
UTC: 2017-10-11 09:35:39

~ Ω timez US/Eastern Asia/Shanghai
US/Eastern: 2017-10-11 05:36:59
Asia/Shanghai: 2017-10-11 17:36:59

~ Ω timez ET central mt US/Pacific
US/Eastern: 2017-10-11 05:38:11
US/Central: 2017-10-11 04:38:11
US/Mountain: 2017-10-11 03:38:11
US/Pacific: 2017-10-11 02:38:11
```

That's all I've had time for this evening - setting default zones, parsing timezone strings, and setting some aliases.
My next steps are documented in the repo's readme.

### Quick Pattern - DI
If `timez` does not see a timestamp in its CLI args, it assumes you're asking about the current time.
My first testcase is for this scenario,  where you pass no arguments:
```go
type testcase struct {
	name     string
	input    string
	expected string
}
testcase{"empty input", "",
`Pacific/Auckland: 2017-10-10 23:01:30
UTC: 2017-10-10 10:01:30
`},
```
There's a problem - I can't control what time you run the tests, so `time.Now()` is going to return something different.
The solution is "dependency injection:" use different code to decide what time it is, depending on whether it's a test or the  actual binary.

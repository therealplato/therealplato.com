---
title: "#soberoctober day 11: CLI timezone conversion"
date: 2017-10-11T22:12:35+13:00
draft: false
---

Movio has offices in US, UK, DE, NZ and has clients globally. I frequently find myself asking "Is it office hours in the US? Is it the middle
of the night in Asia?" We keep records with UTC timestamps (as all programmers should!) but it's still difficult to conceptualize and
mentally convert timezones. 

In the past, I've used several helpful tools:

* The excellent [timeanddate.com converter](https://www.timeanddate.com/worldclock/converter.html?iso=20171010T100100&p1=22&p2=tz_pt) gives you multiple zones and the
ability to type or drag to set any zone to any time.
* I've used my Android 6.x phone as a reference, because its stock clock app supports multiple timezones.
* OSX does not support multiple clocks, wtf? I have recently installed the third party "Hour - World Clock" from App Store, which gives me a multi-clock dropdown from my OSX taskbar.

I like the portability and scriptability of CLI solutions so I have started work on a tool to answer my timezone questions.

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

That's all I've had time for this evening - setting Pacific/Auckland as a default, parsing timezone strings you pass, and setting some
aliases.  My next steps are documented in the repo's [readme](https://github.com/therealplato/timez).

<a id="DI"></a>
### Quick Pattern - Dependency Injection
If `timez` does not see a timestamp in its CLI args, it assumes you're asking about the current time.
My first testcase is for such a scenario:
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
There's a problem - I can't control the time or timezone when you run the tests. `time.Now()` is going to return something that does not
match my `testcase.expected`. The solution is probably the most valuable programming skill I know: "dependency injection." 

DI means that I pass code responsible for child behaviors into the function responsible for parent behaviors, instead of the parent
implementing all behaviors. My parent behavior is "if there are no arguments, return the local timezone and current time."
My child behaviors are "determine the local timezone" and "determine the current time".

You can also inject state variables, but I'm not using any here.

```go
func timez(c clocker, z zoner, args []string) string {
  t0 = c.Now()
  z0 = z.Zone()
  return z0+t0   // simplified
```

If my main function is executing, I inject the "real" dependencies into `timez` :
```go
func main() {
  c := &clock{}
  z := &zone{}
  timez(c, z, args)
}
```

But when I'm testing `timez`, I pass it a mock clock and timezone, whose values I can control:
```go
z0, err := time.LoadLocation("Pacific/Auckland")
z := &mockZone{}
z.On("Zone").Return(z0)

t0, err := time.Parse("2006-01-02 15:04:05-0700", "2017-10-10 23:01:30+1300")
c := &mockClock{}
c.On("Now").Return(t0)

out := timez(c, z, args)
```
Passing these different types to the same functions works thanks to Go's interfaces. In the `timez` function signature you saw it accepts `clocker`:
```go
type clocker interface {
	Now() time.Time
}
```
Both my `clock` and `mockClock` types have a method `Now()` returning a `time.Time`. Thus they both "fulfil the `clocker` interface" and are
acceptable inputs to the `timez` function. You can see both `clock` and `mockClock` implementations in
[`clock.go`](https://github.com/therealplato/timez/blob/6c25bf2b3d2d571b353d8c106605ec933f12f585/clock.go).

---
title: "#soberoctober day 12: finishing timez"
date: 2017-10-12T23:18:07+13:00
draft: false
---

Nothing special to report, I spent the evening improving yesterday's [`timez` project](https://github.com/therealplato/timez). 

I split up my complicated `parse()` function, added many testcases, added some new aliases and formatting strings, loaded local timezones
from ~/.timezrc or `date` binary, and added docs.

I was tickled to see it work across DST boundaries:

```
content/posts Ω timez UTC at 2017-09-24 02:59:00 Pacific/Auckland
UTC: 2017-09-23 14:59:00

content/posts Ω timez UTC at 2017-09-24 03:01:00 Pacific/Auckland
UTC: 2017-09-23 14:01:00
```

I did it for me, but I'd love to hear if you also find [`github.com/therealplato/timez`](https://github.com/therealplato/timez) useful!

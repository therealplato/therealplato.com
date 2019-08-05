---
title: "Breaking Production"
date: 2019-08-06T00:00:00Z
draft: true
---
```
package main
func main(){}

func (i code) am() string {
  return "kind of a big deal"
}

// prog.go:4:9: undefined: code
```

Software falls into two general categories: "Works", "Doesn't Work"; or so we might generalize.

I prefer to think in terms of risk. Risk is a squirmy unit, hard to pin down to numbers. Is the risk of a meteor taking out the server 1% a day? Certainly not! 0.0001% a day? Probably still too high, but who knows? Searching "meteor risk calculator" doesn't help.

What I _do_ evaluate is relative risks. On a given day, a meteor taking out the server is less likely than me deploying code that doesn't work. I should spend my resources on reducing the latter risk, not bombproofing the data center.

The code at the top doesn't compile. We can try and deploy it but it won't work. "Deploying code" means "making it run", so we literally cannot deploy this code.
This is the best kind of failure - one that's impossible to miss. We easily evaluate its failure risk as 100%.

The real interesting risk evaluations happen the rest of the time, when "things might go wrong but probably not."
We have to answer, as best we can:
- Which is it, "might" or "probably?"
- When you say "go wrong", do you mean "it's a bit slow" or "we broke DNS for the western hemisphere?"

Engineers aim to optimize tradeoffs, here, minimizing breakage as we produce correctly behaving code.

Some approaches I've encountered:
- Code nothing: You can't break anything
- Don't minimize breakage: Maybe you get follow-on work to fix your bugs
- Minimize impact: Write code that breaks gracefully
- Minimize risk: Be really sure that it won't break

We have two things to optimize for: minimizing breakage, and producing product code. "Code nothing" minimizes breakage but produces no product code. "Minimize impact" and "minimize risk" give us both reduced breakage and some code.

---

TODO body
Risk analysis of top level vendoring vs service level vendoring vs no vendoring
  Factors in risk calculus
    Scope of change:
      How much code is affected by a change
    Potential negative impact:
      Best case: service stops building
      Worst case: service builds but now uncaught bug spends other people's money
    Potential positive impact:
      identify code that relied on inconsistencies
      identify tests that were insufficient
      plug security holes from outdated dependencies
      guarantee future code consistency
      trust `library.methodCall` means the same thing everywhere
      codebase becomes simpler
    Mitigations:
      testing
      gradual rollout

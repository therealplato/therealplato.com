---
title: "Breaking Production"
date: 2019-08-07T00:00:00Z
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

A hasty generalizer might categorize software into two buckets: "Works", and "Doesn't Work".  They've missed an exciting middle ground: "Works at first, then breaks."

This post discusses some risk assessments and impact mitigations that engineers can use to minimize the costs of releasing broken code.

####Risk

The code above doesn't compile and thus doesn't run. "Deploying code" includes "making it run", so we cannot successfully deploy this code.
This is my favorite kind of failure - one that's impossible to miss. We easily evaluate its failure risk as 100%.

We got lucky there. Risk is often not possible to quantify. Is the risk of a meteor taking out the server 1% a day? Certainly not! 0.000001% a day? Probably still too high, but who knows? Searching "meteor risk calculator" finds global aggregates but no obvious way to evaluate impact risk at a particular location.

In many cases, the best we can do is evaluate relative risks. On a given day, a meteor taking out the server is less likely than me deploying broken code. If my objective is to minimize the risk of the code not running, bombproofing the data center is an option, but my time is probably better spent writing better code.

An overeager and undertrained engineer might state their guess as a quantified risk. A seminal case study of risk management is the series
of 6 massive radiation overdoses, 3 fatal, caused by software bugs in the [Therac-25 ]() radiation therapy machine. After the third
incident, the manufacturer AECL implemented a hardware fix, and claimed that "Analysis of the hazard rate resulting from these modifications
indicates an improvement of at least five orders of magnitude." Without good software design and a comprehensive test suite, they did not
identify the causes of the failures, fixed the wrong thing, and patients continued to die.

The more experienced engineer may still state their guess, but is careful to frame it with language that clearly explains their assumptions
and limitations of scope. A more precise framing would have been "A risk analysis of the
modifications shows five orders of magnitude lower risk of faulty readouts of the turntable position microswitches." It's hard to come up
with even hypothetical math that would explain where the "five orders of magnitude" number came from, so it's entirely possible that this
revised statement is inaccurate.

Risk assessment starts with enumerating assumptions. This matters because the risk model frequently changes if an assumption turns out to be
wrong. The authors of the Therac-25 code assesed software failure risks to be zero. They were (incorrectly) certain that any failures must
be hardware, and their risk assessment was based on this assumption. They were so confident that they removed hardware interlocks that the
previous Therac-20 model used to ensure correct operation. As it turned out, the same software bugs were present in the Therac-20 code, but
did not surface because the interlocks prevented failures. AECL should have assessed a nonzero chance that their assumption of correct code
was faulty.

####Impact

Evaluating failure _impact_ is always specific to the business domain. When you say "failure", do you mean "the page loads
slowly", "we broke DNS for the western hemisphere" or "we killed the patient?" Some domains can accept a higher failure rate than others.


The real interesting risk evaluations happen the rest of the time, when things only _might_ go wrong.

We have to answer, as best we can:
- Which is it, "might" or "probably?"

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

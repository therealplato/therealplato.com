---
title: "Breaking Production"
date: 2019-08-07T00:00:00Z
draft: true
---

# Breaking Production: Risk, Impact and Mitigations

A hasty generalizer might categorize software into two buckets: "Works", and "Doesn't Work". They've missed an exciting middle ground:
"Works until it doesn't."

```go
package main
func main(){}

func (i code) am() string {
  return "kind of a big deal"
}

// prog.go:4:9: undefined: code
```

#### Risk

The code above doesn't compile and thus doesn't run. "Deploying code" includes "making it run", so we cannot successfully deploy this code.
This is my favorite kind of failure - one that's impossible to miss. We easily evaluate its failure risk as 100%.

We got lucky there. Risk is often not possible to quantify. Is the risk of a meteor taking out the server 1% a day? Certainly not! 0.000001%
a day? Probably still too high, but who knows? Searching "meteor risk calculator" finds global aggregates but no obvious way to assess
impact risk at a particular location.

Often the best we can do is evaluate relative risks. On a given day, a meteor taking out the server is less likely than me deploying broken
code. If my objective is to maximize uptime, bombproofing the data center is an option, but my time is probably better spent writing better
code.

An overeager engineer might state a guess as a quantified risk. A seminal case study of risk management is the series of 6 massive radiation
overdoses, 3 fatal, caused by software bugs in the [Therac-25 ](/misc/therac-25.pdf)[^1] radiation therapy machine. After the third incident,
the manufacturer AECL implemented a hardware fix, and claimed that "Analysis of the hazard rate resulting from these modifications indicates
an improvement of at least five orders of magnitude." 

AECL didn't show their math, and it's hard to imagine how they reached this figure. It's safe to say that their evaluation included a faulty
assumption that the software worked as intended. 

Risk assessment requires enumerating assumptions and considering the possibility that those assumptions are faulty. This matters because the
risk model frequently changes when an assumption turns out to be wrong. The authors of the Therac-25 code were certain that any failures
must be hardware, and assumed that the software failure risk was zero. They were so confident that they removed hardware interlocks used in
the previous Therac-20 model. As it turned out, the same software bugs were present in the Therac-20 code, but did not surface because the
interlocks prevented failures. AECL's assessment should have assigned a nonzero probability to the risk that their assumption was wrong.

Given that risk assessments involve unknowns, sometimes a guess is the best we can do. It's important to provide context that clearly
communicate any assumptions and limitations of scope. A more precise framing of AECL's statement would have been "A risk analysis of the
modifications shows five orders of magnitude lower risk of failure caused by incorrect readouts from the turntable position microswitches."
Users trusted AECL's assessment, but the actual causes remained unidentified. AECL fixed the wrong thing, and patients continued to die.

#### Impact

Evaluating _failure impact_ is always specific to the business domain. When you say "failure", do you mean "the page doesn't load", "we owe lots of money" or "we killed the patient?" Some domains can tolerate failure more than others.

We can also assess _change impact_: What is the set of things that this change could break? The trivial case is changing a process that runs
in total isolation. This change touches affects that process, so only that process could fail. The ultimate case is a change across the
board, perhaps rewriting the whole system in a new language or replacing an omnipresent storage system. This is a high-impact change because
it touches everything; if something goes wrong, everything could fail. As you'd expect, most changes fall in between these extremes.

Factors worth considering when assessing impact include:
- Physical harm
- Person-hours spent
- Monetary cost
- Reputation loss

As with evaluating risks, it's important to also consider how the impact changes if it turns out your assumptions are faulty.

#### Risk Calculus
I've used the term _risk calculus_ before but never formally defined it to myself. Upon reflection it's a tongue-in-cheek term because
mathematical calculus is exact but risk is fuzzy. Let's give it a go:

```
    Failure Mode is defined as "An Incorrect Behavior"
    Risk is defined as "Probability of a Failure Mode occurring"
    Impact is defined as "Cost of a Failure Mode"

    Do not use the system until:

    The sum of (Risk * Impact) for all Failure Modes 
      is less than
    The benefit when zero Failure Modes occur
      plus
    The sum of (the benefit of each Failure Mode) for all Failure Modes
```

The benefit of a failure mode is usually limited to "we learned this failure mode exists, how to detect it and how to avoid it." In
low-impact domains, the value of this learning experience often exceeds the costs of failure.[^2]

It's up to the business to put numbers on the costs and benefits, or at least to determine relative weights. The engineers building the
system should be willing to explain their own risk calculus if they disagree with the business's calculus. This includes scenarios where the
engineer identifies an ethical cost that was overlooked or undervalued by the business.

It's also worth reiterating that perfect information is almost never available. Assume the worst case.

#### Mitigations
without good software design and a comprehensive test suite, 
Engineers aim to optimize tradeoffs, here, minimizing breakage as we produce correctly behaving code.
Resilience: Code independent pieces that are unaffected by changes to other parts of the system

Some approaches I've seen:
- Code nothing: You can't break anything
- Code poorly: Your government contract includes follow-on work to fix your bugs
- Minimize impact: Write code that breaks gracefully
- Minimize risk: Be really sure that it won't break

We have two things to optimize for: minimizing breakage, and producing product code. "Code nothing" minimizes breakage but produces no product code. "Minimize impact" and "minimize risk" give us both reduced breakage and some code.

[^1]:
"Medical Devices: The Therac-25". Nancy Leveson, University of Washington. Via [danluu/post-mortems](https://github.com/danluu/post-mortems)

[^2]:
"Move fast and break things." Mark Zuckerberg, [Facebook IPO filing](https://www.sec.gov/Archives/edgar/data/1326801/000119312512034517/d287954ds1.htm#toc287954_10), 2012

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
      defense in depth


The real interesting risk evaluations happen the rest of the time, when things only _might_ go wrong.

We have to answer, as best we can:
- Which is it, "might" or "probably?"


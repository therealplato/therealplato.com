---
title: "Breaking Production: Risk, Impact and Mitigations"
date: 2019-08-12T00:00:00Z
draft: false
---


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

Often the best we can do is evaluate relative risks - "it might happen" vs "it probably will happen." On a given day, a meteor taking out the
server is less likely than me deploying problematic code. If my objective is to maximize uptime, bombproofing the data center is an option,
but my time is probably better spent writing better code.

An overeager or pressured engineer might state a guess as a quantified risk. An excellent [case study](/misc/therac-25.pdf)[^1] of risk
management is the series of 6 massive radiation overdoses, 3 fatal, caused by software bugs in the Therac-25 radiation therapy machine.
After the third incident, the manufacturer AECL implemented a hardware fix, and claimed that "Analysis of the hazard rate resulting from
these modifications indicates an improvement of at least five orders of magnitude."

AECL didn't show their math, and it's hard to imagine how they reached this figure. It's safe to say that their evaluation included a faulty
assumption that the software worked as intended.

Risk assessment requires enumerating assumptions and considering the possibility that those assumptions are faulty. This is important
because the risk model frequently changes when an assumption turns out to be wrong. AECL's assessment should have considered the fatal
consequences of their incorrect assumption.

Given that risk assessments involve unknowns, sometimes a guess is the best we can do. It's important to provide context that clearly
communicates any assumptions and limitations of scope. AECL could have more precisely framed their statement by saying "A risk analysis of
the modifications shows five orders of magnitude lower risk of failure caused by incorrect readouts from the turntable position
microswitches." The actual causes remained unidentified, AECL fixed the wrong thing, and patients continued to die.

#### Impact

Evaluating _failure impact_ is always specific to the business domain. When we say "failure", do we mean "the page doesn't load", "we owe
lots of money" or "we killed the patient?" Some domains can tolerate failure more than others.

We can also assess _change impact_: What is the set of things that this change could break? The trivial case is changing a process that runs
in total isolation - only that process can fail. The ultimate case is a change across the board, perhaps rewriting the whole system in a new
language or replacing an omnipresent storage system. These are high-impact changes because they touch everything; if something goes wrong,
everything can fail. Most changes fall in between these extremes.

Factors worth considering when assessing impact include:
- Unethical system behavior, including physical harm
- Person-hour costs
- Monetary costs
- Reputation costs

As with evaluating risks, it's important to also consider how the impact would change if any assumptions are incorrect.

#### Risk Calculus

I've used the term _risk calculus_ in conversations but never formally defined it to myself. Upon reflection it's a tongue-in-cheek term because
mathematical calculus is exact but risk is fuzzy. Let's give it a go:

```
    Failure is defined as "An Incorrect Behavior"
    Risk is defined as "Probability of the Failure occurring"
    Impact is defined as "Cost of the Failure"

    Do not use the system until:

    The sum of (Risk * Impact) for all Failures
      is less than
    (The probability of no Failures * The benefit of no Failures)
      plus
    The sum of (the benefit of the Failure) for all Failures
```

Benefits of failure include learning about an unknown failure mode, learning how to detect it, learning how to avoid it, identifying missing
tests, and identifying previously unknown assumptions. The value of this learning experience often exceeds the costs of failure.[^2]

It's up to the business to put numbers on the costs and benefits, or at least to assign relative weights. The engineers building and
operating the system should be willing to explain their own risk calculus when it differs from the business's calculus. This notably
includes scenarios where an engineer identifies an ethical impact that was overlooked or undervalued by the business.

It's also worth reiterating that perfect information is almost never available. Weighing risks and impacts higher than reality is safer than
weighing them too low. Recalculate when new information becomes available.

#### Mitigations

Engineering is all about optimizing tradeoffs. In the context of risk mitigation, I've seen several approaches:
- Code nothing: We can't break anything.
- Code poorly: We'll get a follow-on contract to fix our bugs.
- Minimize impact: Write code that breaks gracefully.
- Minimize risk: Write code that has very low chance of breaking.
- Maximize benefit: Write code that is so valuable that users tolerate breakage.

Personally, I also value the engineering team's happiness, so I prefer a combination of minimizing impact and minimizing risk.

##### Minimize Impact

Build systems that handle known failure modes, but let the engineers handle unknown failure modes:
- Build high-availability systems with multiple replicas to minimize impact when networking or hardware problems make a service unavailable.
- Retry after transient failures (but not indefinitely.)
- Retry if dependencies are unavailable at system startup (but not indefinitely.)
- Do not take zero action when an error or exception is encountered.
- Prefer terminating the process over continuing with an inconsistent or unknown program state.
- Use instrumentation and alerting to quickly identify failures.

##### Minimize risk

Code defensively. Assume the worst and try to code around it:
- Do not blindly assume that your code behaves as expected. Write validation or tests to confirm the behavior. Don't be afraid of checking assumptions repeatedly.
- Do not assume that third party code behaves as expected.
- Make services as independent as reasonably possible to minimize change risk. For instance, instead of having two services share a database
  table, give one service responsibility for maintaining that table and expose an API through which other services can use the table.
- Gradually roll out changes that touch multiple services, rather than all at once.
- Have a comprehensive test suite, including validating the system's behavior when errors occur or assumptions are faulty.
- Do not use concurrency when synchronous code will do the job.
- Do not use ten layers of indirection when one layer will do the job.

The manufacturers of the Therac-25 code incorrectly calculated their risks. They were so confident in their software that they removed
hardware interlocks used in the previous Therac-20 model. As it turned out, the same bugs were present in the Therac-20 as the Therac-25,
but were not noticed because the interlocks prevented accidents. Do not mistake absence of evidence for evidence of absence.

---

[^1]:
1: "Medical Devices: The Therac-25". Nancy Leveson, University of Washington. Via [danluu/post-mortems](https://github.com/danluu/post-mortems)

[^2]:
2: "Move fast and break things." Mark Zuckerberg, [Facebook IPO filing](https://www.sec.gov/Archives/edgar/data/1326801/000119312512034517/d287954ds1.htm#toc287954_10), 2012

---
title: "#soberoctober day 18: coding rules"
date: 2017-10-18T20:37:25+13:00
draft: false
---

Tonight I've begun formalizing my definition from [day 15](/posts/soberoctober15), continuing to work towards the goal of programmatic rules.

> N.B: In this post I use the word "behavior" to apply to both code (where it means: business logic, described by a function signature, implemented by code) and also to humans (where it means: things humans do.) Sorry for any confusion.

A key behavior of a rule is _does it apply in this scenario?_
```
type Rule interface {
	Governs(subject subverse.Identity, context string) bool
}
```
I plan to extend this interface. Another behavior that I am certainly missing is:

* Is the rule broken if the subject does behavior X?

That's pretty hard to answer. In some circumstances, a computer can answer unaided: _The CI failed, therefore merge is blocked._ Other times
a human will have to decide: _This art is not pornographic, therefore it may remain on the site._

Deciding "Which human adjudicates" is an important question and can itself be governed by rules. Some of today's counterparties already do
this - the majority of those EULA's you skip through have a clause that say "You agree that disputes will be resolved by arbitration service
XYZ instead of a civil court." 

The _context wherein a rule applies_ is a wide question and may get squirrely:

* **Men** must use the men's restroom. (But I self identify as...!)
* **Customers** must... (But I'm not buying anything!)
* **Citizens** must... (But I'm rich and powerful!)
* **Players of this tournament** must...
* **Sunnis** must... **Catholics** must... **Shiites** must... **Scientologists** must... **Jews** must...

I've coded the bare basics:
```
// rule.go
type SimpleRule struct {
	Ruler       subverse.Identity
	Subject     subverse.Identity
	Behavior    string
	Context     string
	Consequence string
}
...

// rule_test.go
package rule

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/therealplato/subverse"
)

func TestSimpleRuleDoesNotGovernAnonymous(t *testing.T) {
	soberoctober := SimpleRule{
		Ruler:       subverse.Anonymous,
		Subject:     subverse.Anonymous,
		Context:     "october 2017",
		Behavior:    "Do not consume psychoactives, except caffeine. Do not play video games.",
		Consequence: "feel shame",
	}

	assert.False(t, soberoctober.Governs(subverse.Anonymous, "october 2017"))
}

func TestSimpleRuleGovernsSubjectWithinContext(t *testing.T) {
	me := &subverse.NamedIdentity{"plato"}
	subject := me
	soberoctober := SimpleRule{
		Subject: subject,
		Context: "october 2017",
	}

	assert.True(t, soberoctober.Governs(me, "october 2017"))
}

func TestSimpleRuleDoesNotGovernSubjectOutsideContext(t *testing.T) {
	me := &subverse.NamedIdentity{"plato"}
	subject := me
	soberoctober := SimpleRule{
		Subject: subject,
		Context: "november 2017",
	}

	assert.False(t, soberoctober.Governs(me, "october 2017"))
}

func TestSimpleRuleDoesNotGovernNonSubjects(t *testing.T) {
	me := &subverse.NamedIdentity{"plato"}
	soberoctober := SimpleRule{
		Subject: &subverse.NamedIdentity{"Cat."},
		Context: "october 2017",
	}

	assert.False(t, soberoctober.Governs(me, "october 2017"))
}
```
I really like test-driven design. It forces me to use the code before I write the code. Sometimes the test is so hard to write that I realize my usage pattern (aka the code's API) makes no sense. Great! That means I can throw it out, jump ahead and design a better API - without wasting any time actually implementing the crappy API.

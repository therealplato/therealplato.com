---
title: "#soberoctober day 26: playmat plans"
date: 2017-10-26T23:36:49+13:00
draft: false
---

I mentioned Inked Gaming in [#soberoctober25](/posts/soberoctober25) and their [custom playmat service](https://inkedgaming.com).

I've come up with an design for my own playmat. Or at least, the plan for a design; if my past forays into digital art are any indication, there
will be a lot of sweat and tears before I've got a finished file that I'm happy with.

The plan:

* Take SVG images of the [Go Gopher](https://github.com/egonelbre/gophers).
* Tint the gophers with the five colors of MTG mana.  
* Ideally: bling the gophers out with things that represent those colors - a skull for black, fire for red, elf ears for green, [sunglasses](/images/dwi.gif) for blue, angel wings for white
* Place the gophers in a circle around the middle of the playmat.  
* Draw lines connecting various color pairs/triads.
* Label the lines with the [names of the corresponding tribes](https://boardgames.stackexchange.com/questions/11550/what-are-the-names-for-magics-different-colour-combinations/11563).  
* Off to the side, place this code to reference parts of a turn:

```go
func turn() {
  beginning: {
    untap()
    upkeep()
    draw()
  }
  main1: {
    if !landThisTurn {
      playLand()
    }
    castSpells()
  }
  combat: {
    begin()
    declareAttackers()
    declareBlockers()
    combatDamage()
    end()
  }
  main2: {
    if !landThisTurn {
      playLand()
    }
    castSpells()
  }

  end: {
    endStep()
    cleanupDmg()
  }
}
```

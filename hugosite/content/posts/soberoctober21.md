---
title: "#soberoctober day 21: combinatorics beat me"
date: 2017-10-21T23:50:52+13:00
draft: false
---

## fun:
I'm still having a real blast playing Magic downtown at [King Of Cards](http://www.yourlocalgameshop.com/). I won about $6 worth of cards
after entering a free "Standard Showdown" tournament with a borrowed deck. Neat.

After the shop closed, I spent a good hour looking at all the custom playmats available at [Inked Gaming](https://inkedgaming.com). There's
a bunch of nice ones but I want to design one that has the Golang gopher on it. Playmats are just soft mats you put your cards on. They
can protect cards from the table surface, but since I use card sleeves, I'd only be buying one for the looks.

## more fun:
**@pltvs** from the Gophers slack linked a [Go challenge](http://pliutau.com/practice-go-secret-message/). My solution used the same basic approach as other submissions but mine's [a bit slower](https://github.com/plutov/practice-go/pull/33/commits/f56f032edb0a203c85093b07550a7bd285378990).
## aaargh:
After pull-requesting that submission, I spotted a failed travis build and decided to fix it. Then I ate humble pie because the failure stems
from a challenge I can't solve yet. I'll [paraphrase it](https://github.com/plutov/practice-go/tree/master/coins) as:

> Given _n_ coins, separate them into _m_ piles of size _m1_, _m2_...  
> What is the size of the set of pile sizes?  
> For example, the set of pile sizes for:  
> _n=3_ is _(3) (2 1) (1 1 1)_  
> _n=4_ is _(4) (3 1) (2 2) (2 1 1) (1 1 1 1)_  
> and the size of those sets are 3 and 5 respectively.  

If you've got a solution, [populate this function](https://github.com/plutov/practice-go/blob/master/coins/coins.go), run the tests, and
open a PR!

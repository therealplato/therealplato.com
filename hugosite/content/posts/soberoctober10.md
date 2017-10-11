---
title: "#soberoctober day 10: svg ftw (edit: not yet)"
date: 2017-10-10T21:27:08+13:00
draft: false
---

Long day today, leaving me with no time to [dive deep](https://www.youtube.com/results?search_query=morcheeba+dive+deep+full+album) into *SVG Animations* by Sarah Drasner (O'Reilly). Copyright 2017 Sarah Drasner, 978-1-491-93970-3.

Let's see what I can get done with an hour and some lipsum...

<hr>

... Not very far. My objective was to make an arrow that starts pointing down and animates to point up as you scroll
down the page. 

After conceding that was not a newbie task, I just looped the animation.

I don't see a way to set an infinite repeat count with the `animation:` css shorthand.

 I first tried animating the `x2, y2` properties of the arrowhead line segments, but
apparently that's not a thing and I need to use `transform` CSS property or `transform` property of an SVG element.
I did not find any examples of animating the latter.

As you can see, I haven't figured out how to chain `rotate` and `translate`.

I did learn about the `viewBox` property and have used it to place my SVG's origin at the center, which made a huge
improvement.

Almost all docs I've found give up and use a JS library. It would be nice to avoid that.

Maybe I'll learn a correct way to do it and update this post?

<hr>

<script>
</script>
<style>
#flipper-container{
	position: fixed;
  top: 0%;
  left: 50px;
  perspective: 600px;
  -webkit-perspective: 600px;
  -moz-perspective: 600px;
  animation: scroll-animation 2s 2s 999 alternate ease-in-out forwards;
}
#flipper {
  width: 100%;
  height: 100%;
  font-size: 80px;
  line-height: 100px;
  transform-style: preserve-3d;
  animation: flip-animation 2s 2s 999 alternate ease-in-out forwards;
}
.face {
  position: absolute;
  backface-visibility: hidden;
}
#back {
  transform: rotateY( 180deg );
}
@keyframes scroll-animation {
  0% {
    top: 0%;
  }
  100% {
    top: 100%;
  }
}
@keyframes flip-animation {
  0% {
    transform: rotateX(0deg);
  }
  100% {
    transform: rotateX(180deg);
  }
}
#slider{
	position: fixed;
  top: 80px;
  right: 50px;
	width: 80px;
	height: 80px;
}
.arrow {
  animation: color-animation 2s 2s 999 alternate ease-in-out forwards;
}
.head-left {
  animation: head-left 2s 2s 999 alternate ease-in-out forwards;
}
@keyframes color-animation {
  0% {
    stroke: black;
  }
  100% {
    stroke: red;
  }
}
@keyframes head-left {
  0% {
    transform:rotate(45deg) translateY(40px) ; 
  }
  100% {
    transform:rotate(-45deg) translateY(-40px) ;
  }
}
</style>
<svg id="slider" viewBox='-40 -40 80 80'>
  <g class="arrow" fill="none" stroke="black">
		<line x1="0" y1="-40" x2="0" y2="40" />
		<line class="head-left" x1="0" y1="0" x2="-20" y2="0" />
		<line class="head-right" x1="0" y1="40" x2="20" y2="20" />
	</g>
</svg>
<div id="flipper-container">
  <div id="flipper">
    <span id="front" class="face">↓</span>
    <span id="back" class="face">↓</span>
  </div>
</div>

Leverage agile frameworks to provide a robust synopsis for high level overviews. Iterative approaches to corporate strategy foster collaborative thinking to further the overall value proposition. Organically grow the holistic world view of disruptive innovation via workplace diversity and empowerment.

These old Doomsday Devices are dangerously unstable. I'll rest easier not knowing where they are. This opera's as lousy as it is brilliant! Your lyrics lack subtlety. You can't just have your characters announce how they feel. That makes me feel angry!

Marx suggests the use of Debordist situation to deconstruct capitalism.
Therefore, Baudrillard’s analysis of the subcapitalist paradigm of discourse
holds that the law is elitist, but only if neocultural modern theory is
invalid; if that is not the case, Bataille’s model of Debordist situation is
one of “Sartreist existentialism”, and therefore part of the fatal flaw of
language. 

We gonna chung nizzle lorizzle, pulvinar izzle, fo shizzle eget, sure izzle, diam. Etizzle sizzle leo its fo rizzle sem hendrerizzle mattis. Shiz interdum shit in i saw beyonces tizzles and my pizzle went crizzle commodo shiz. Etiam pizzle fermentizzle ligula. Morbi fo shizzle mah nizzle fo rizzle, mah home g-dizzle. Maecenas quis metizzle ac dolor iaculizzle auctor. Shut the shizzle up sagittizzle viverra urna. Shiznit sollicitudizzle massa fo the bizzle. Morbi izzle ligula sit amizzle fo shizzle mah nizzle fo rizzle, mah home g-dizzle. Nam you son of a bizzle enizzle vitae doggy. Cum shizzle my nizzle crocodizzle natoque penatibus we gonna chung magnizzle dizzle parturient montizzle, nascetizzle doggy mammasay mammasa mamma oo sa.

Several discourses concerning cultural desituationism exist. Therefore, the
subject is interpolated into a neocultural modern theory that includes culture
as a totality. 

But existing is basically all I do! Yep, I remember. They came in last at the Olympics, then retired to promote alcoholic beverages! But I've never been to the moon! Switzerland is small and neutral! We are more like Germany, ambitious and misunderstood!

Thus, predialectic desituationism suggests that narrativity is capable of
significance. An abundance of sublimations concerning Lacanist obscurity exist.

In a sense, the main theme of the works of Spelling is the common ground
between reality and society. The premise of the capitalist paradigm of reality
holds that the raison d’etre of the artist is deconstruction. 

Anyone who laughs is a communist! We can't compete with Mom! Her company is big and evil! Ours is small and neutral! Why, those are the Grunka-Lunkas! They work here in the Slurm factory. Noooooo!

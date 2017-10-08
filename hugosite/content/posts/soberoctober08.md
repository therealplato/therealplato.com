---
title: "#soberoctober day 8: interplanetary file system"
date: 2017-10-08T21:17:55+13:00
draft: false
---

This evening I set up the Interplanetary File System and uploaded the first episode of Reality Exploit Roundtable, a 2012 podcast I helped
create.

It was really straightforward:
```sh
go get -u -d github.com/ipfs/go-ipfs
cd $GOPATH/src/github.com/ipfs/go-ipfs
make install
ipfs init
ipfs daemon #in other term
ipfs add ~/Downloads/realityexploitroundtable-EP001.ogg
# added QmdPt2EqonqKcwyxbuePuiwPsEc8c3t5R3Dr4oTDBpUmXM realityexploitroundtable-EP001.ogg
```
Tada! That's that, my file is now in a global p2p filesystem. Actually, "global" is a misnomer - it's named "interplanetary" because the
project aims to be usable for syncing content between Earth and Mars and beyond, and is designed to work within bandwidth and latency
constraints.

Several members of the IPFS community offer HTTP gateways to IPFS content hashes. Have a listen if you'd like to hear me and my colleagues
discuss privacy pre-Snowden. ~40 minutes
<br><a href="https://ipfs.io/ipfs/QmdPt2EqonqKcwyxbuePuiwPsEc8c3t5R3Dr4oTDBpUmXM">https://ipfs.io/ipfs/QmdPt2EqonqKcwyxbuePuiwPsEc8c3t5R3Dr4oTDBpUmXM</a>
<br><a href="https://ipfs.infura.io/ipfs/QmdPt2EqonqKcwyxbuePuiwPsEc8c3t5R3Dr4oTDBpUmXMk">https://ipfs.infura.io/ipfs/QmdPt2EqonqKcwyxbuePuiwPsEc8c3t5R3Dr4oTDBpUmXMk</a>

If you have IPFS installed yourself, you can retrieve the podcast with:
```sh
ipfs get QmdPt2EqonqKcwyxbuePuiwPsEc8c3t5R3Dr4oTDBpUmXM -o=realityexploitroundtable-EP001.ogg
```

I gotta give a shoutout to the IPFS contributors for their ascii art logo. It looks outstanding in cool-retro-term:
<image src="/images/crt-ipfs.png"/>

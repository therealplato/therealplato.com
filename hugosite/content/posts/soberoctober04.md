---
title: "#soberoctober day 4: a bit of serverless"
date: 2017-10-04T19:34:11+13:00
draft: false
---

## Today's Achievements:
I tried Amazon Lambda.  I first heard of this "serverless" computing at Full Stack Fest 2016. You give Amazon a function and only pay for
resources used as the function runs. I immediately observed that there are, in fact, still servers involved, you're just delegating
operations to Amazon.

After I got over my pedantry and listened to [this story](https://www.youtube.com/watch?v=9IrFIobZUEA&t=22m10s) I decided I liked the idea.
Lambda can be orders of magnitudes cheaper than EC2 (for sporadic resource use,) or more expensive than EC2 (for pegged resource use.)

### DONE:
* Created AWS account
* Read Lambda, DynamoDB docs and tutorieals
* Created IAM role, group, user
* Installed and configured awscli
* Created a streaming Dynamo table
* Configured the Dynamo table to communicate with Lambda
* Created `com-therealplato-counter` Lambda function, that receives HTTP and pings Dynamo

### TODO:
Fix the function! The output indicates I'm giving Dynamo bad information:
```
There were 5 validation errors:\n* MissingRequiredParameter:
Missing required key 'TableName' in params\n*
MissingRequiredParameter: Missing required key 'Item' in
params\n* UnexpectedParameter: Unexpected key 'Items' found
in params\n* UnexpectedParameter: Unexpected key 'Count'
found in params\n* UnexpectedParameter: Unexpected key
'ScannedCount' found in params
```

## Just Arch Things...
I haven't yet found an obvious way to make my volume keys adjust volume. Apparently I can edit xorg configuration
files to manually map the keycodes to XF86 audio events.

F Dat! `alsamixer` command works well enough for now.

*thatsthejoke.jpg*: `cool-retro-term` abbreviates to `CRT` 

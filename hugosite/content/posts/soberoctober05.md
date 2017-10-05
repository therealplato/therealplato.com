---
title: "#soberoctober day 5 - JIT"
date: 2017-10-05T23:47:00+13:00
draft: true
---
<script type="text/javascript">
document.addEventListener("DOMContentLoaded", function() {
  $.get("https://5b567rkzxe.execute-api.us-east-2.amazonaws.com/prod/com-therealplato-counter?TableName=com-therealplato-counter",
  function(data, status){
    if(status != "success") { return }
		$("#counter-button").text(data.m)
	})

	$("#counter-button").click(function(){
		$.post("https://5b567rkzxe.execute-api.us-east-2.amazonaws.com/prod/com-therealplato-counter?TableName=com-therealplato-counter", {}, function(data, status){
			if(status != "success") { return }
			$("#counter-button").text(data.m)
		});
	}); 
})
</script>

I barely completed this by midnight! AWS is complicated.  This button hits an AWS API Gateway, which proxies to a lambda function, which
reads and/or updates dynamodb:

<span>Button has been clicked</span>&nbsp;<button class="counter" id="counter-button">...</button>&nbsp;<span>times</span>


## Arch Of The Day
 
While trying to use two finger scrolling, I observed this entertaining behavior:
```gherkin
Given my cursor is on a link
When I touch one finger
And then a second finger
And then drag up or down
Then each scroll event opens a new tab
```
This behavior can easily open 50 copies of the link under the cursor.

It appears resolved after disabling some `middlemouse.*` settings in firefox's `about:config`.

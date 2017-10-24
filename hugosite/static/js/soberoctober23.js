var url = "/misc/suzuka-laps.json";
d3.json(url, function (data) {
  console.log(data)
  var bars = d3.select("#laptimes-container")
    .selectAll("span")
    .data(data);

  var tmax = d3.max(data, function(d, i) {
    return d.timeMilliseconds
  })
  var tmin = d3.min(data, function(d, i) {
    return d.timeMilliseconds
  })

	var xScale = d3.scaleLinear()
			.domain([0, tmax]) // lap times
			.range([0, 100]); // percent widths 

  bars.enter()
      .append("span")
      .style("display", "block")
			.style("margin-bottom", "4px")
      .style("color", "white")
      .style("background-color", "black")
      .style("height", "1em")
      .style("line-height", "normal")
      .style("width", function(d) {
        return xScale(d.timeMilliseconds)+"%";
      })
			.text(function(d){
				return d.year + ": " + d.driver + " " + d.timeDisplay
			});

  bars.exit()
      .remove();


});

var url = "/misc/2001.alonso.laps.json";
d3.json(url, function (rawData) {
  var data = process(rawData);

  // var tmax = d3.max(data, function(d) {
  //   return d.ms;
  // });
  // var tmin = d3.min(data, function(d){
  //   return d.ms;
  // });
  var tmax = Atoi("1:50.000").ms;
  var tmin = Atoi("1:20.000").ms;

  var margin = {top: 10, right: 10, bottom: 30, left: 60},
        width = 600 - margin.left - margin.right,
        height = 400 - margin.top - margin.bottom;

	var xScale = d3.scaleLinear()
			.domain([0, data.length-1]) // lap number
			.range([0, width]); // percent widths 
  var xAxis = d3.axisBottom(xScale)
      .ticks(10);

	var yScale = d3.scaleLinear()
			.domain([tmin, tmax]) // lap times
			.range([0, height]);
  var yAxis = d3.axisLeft(yScale)
      .tickFormat(function(d) {
        return Itoa(d);
      })
      .ticks(10);


  var chart = d3.select("#laptimes-container")
    .append("svg")
    .attr("height", height+margin.top+margin.bottom)
    .attr("width", width+margin.left+margin.right)
  .append("g")
    .attr("transform", "translate(" + margin.left + "," + margin.top + ")");

  chart.append("g")
      .attr("class", "x axis")
      .attr("transform", "translate(0," + height + ")")
      .call(xAxis);

  chart.append("g")
      .attr("class", "y axis")
      .call(yAxis);

  var points = chart
    .selectAll("circle")
    .data(data);

  points.enter()
      .append("circle")
      .attr("cx", function(d, i) {
        return xScale(i)
      })
      .attr("cy", function(d, i) {
        return yScale(d.ms);
      })
      .attr("r", 5);

  points.exit()
      .remove();


});

function process(rawData) {
  var out = rawData.map(Atoi)
  return out;
}

var timing = /(\d+):(\d+)\.(\d+)/

function Atoi(d) {
  var m = d.match(timing);
  if(m == null || m.length != 4){
    return {}
  }
  var ms = +m[3] + (1000 * +m[2]) + (60 * 1000 * +m[1]);
  return {display: d, ms: ms};
}

function Itoa(ms) {
  var mm = Math.floor(ms/60000);
  ms = ms % 60000;
  var ss = Math.floor(ms/1000);
  ms = ms % 1000;
  ms = ("000" + ms).slice(-3)
  return mm+":"+ss+"."+ms;
}

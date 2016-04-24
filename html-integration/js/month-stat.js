$(function () {
   // rearrange month form
   var monthMapping = {
      '1': 'Jan.',
      '2': 'Feb.',
      '3': 'Mar.',
      '4': 'Apr.',
      '5': 'May.',
      '6': 'Jun.',
      '7': 'Jul.',
      '8': 'Aug.',
      '9': 'Sep.',
      '10': 'Oct.',
      '11': 'Nov.',
      '12': 'Dec.'
   }
   var monthXAxis = [];
   var monthYAxis = [];
   var monthYAxisLastYear = [];

   for (var i = 1; i <= 12; i ++) {
      monthXAxis.push(monthMapping[i]);
      monthYAxis.push(dataset["result"]["2015"]["month"][i])
      monthYAxisLastYear.push(dataset["result"]["2014"]["month"][i])
   }

   $('#monthCountChart').highcharts({
      chart: {
         type: 'area'
      },
      title: {
         text: ''
      },
      xAxis: {
         categories: monthXAxis,
      },
      yAxis: {
         title: {
            text: 'Book count'
         },
      },
      tooltip: {
          pointFormat: '<b>{point.y}</b><br/>{series.name}'
         //pointFormat: '{series.name} produced <b>{point.y:,.0f}</b><br/>warheads in {point.x}'
      },
      plotOptions: {
         area: {
            marker: {
               enabled: false,
               symbol: 'circle',
               radius: 2,
               states: {
                  hover: {
                     enabled: true
                  }
               }
            }
         }
      },
      series: [{
         name: '2015',
         data: monthYAxis
      }, {
         name: '2014',
         data: monthYAxisLastYear
      }]
   });
});
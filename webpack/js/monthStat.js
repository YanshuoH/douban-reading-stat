import $ from 'jquery'
import {sortObjectByValue} from './sorting'

let monthMapping = {
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

class MonthStat {
  constructor() {
    this.selector = 'monthCountChart'
    this.top3Container = $('#monthTop3')
  }

  render(currentYearResult, previousYearResult) {
    let monthXAxis = []
    let monthYAxis = []
    let havePreviousYear = previousYearResult == undefined ? false : true
    let monthYAxisLastYear = [];

    for (let i = 1; i <= 12; i ++) {
       monthXAxis.push(monthMapping[i])
       monthYAxis.push(currentYearResult['month'][i])
       if (havePreviousYear) {
         monthYAxisLastYear.push(previousYearResult['month'][i])
       }
    }

    this.chart = new window.Highcharts.Chart(
      this.getHighchartOption(currentYearResult['year'], monthXAxis, monthYAxis, monthYAxisLastYear)
    )

    // Top 3
    this.top3Container.empty()
    let sorted = sortObjectByValue(currentYearResult['month'])
    for (let i = 0; i < 3; i ++) {
      this.top3Container.append(`<li>${monthMapping[sorted[i]]}: ${currentYearResult['month'][sorted[i]]}</li>`)
    }
  }

  getHighchartOption(year, monthXAxis, monthYAxis, monthYAxisLastYear) {
    let options = {
       chart: {
         renderTo: this.selector,
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
          name: year,
          data: monthYAxis
       }]
    }

    if (monthYAxisLastYear.length > 0) {
      options.series.push({
         name: year - 1,
         data: monthYAxisLastYear
      })
    }

    return options
  }

  destroy() {
    this.chart.destroy()
  }
}

export default MonthStat

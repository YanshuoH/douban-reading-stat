import $ from 'jquery'
import {sortObjectByValue, stringRepeat} from './sorting'

class PieChart {
  constructor() {
    this.ratingTop3Container = $('#ratingTop3')
    this.ratingChartSelector = 'ratingChart'

    this.authorTop3Container = $('#authorTop3')
    this.authorChartSelector = 'authorChart'

    this.translatorTop3Container = $('#translatorTop3')
    this.translatorChartSelector = 'translatorChart'

    this.publisherTop3Container = $('#publisherTop3')
    this.publisherChartSelector = 'publisherChart'

    this.Highcharts = window.Highcharts
  }

  // Represent dataset["result"]["2015"]["rating"]
  renderRating(result) {
    var ratingPie = []
    for (var i = 1; i <= 5; i ++) {
      var currentRatingCount = result[i];

      ratingPie.push({
        name: i + " star" + (i == 1 ? "" : "s"),
        y: currentRatingCount
      });
    }

    this.ratingchart = new this.Highcharts.Chart(
      this.getHighchartOption('Rating', ratingPie, this.ratingChartSelector)
    )

    // Top 3
    this.ratingTop3Container.empty()
    let sorted = sortObjectByValue(result)
    for (let i = 0; i < 3; i ++) {
      let starsString = stringRepeat(`<i class="fa fa-star" aria-hidden="true"></i>`, sorted[i])
      this.ratingTop3Container.append(`<li>${starsString}: ${result[sorted[i]]}</li>`)
    }
  }

  // Represent dataset["result"]["2015"]["author"]
  renderAuthor(result) {
    var authorPie = []
    for (var author in result) {
      authorPie.push({
        name: author,
        y: result[author]
      });
    }

    this.authorChart = new this.Highcharts.Chart(
      this.getHighchartOption('Author', authorPie, this.authorChartSelector)
    )

    // Top 3
    this.authorTop3Container.empty()
    let sorted = sortObjectByValue(result)
    for (let i = 0; i < 3; i ++) {
      this.authorTop3Container.append(`<li>${sorted[i]}: ${result[sorted[i]]}</li>`)
    }
  }

  getHighchartOption(seriesName, data, selector) {
    return {
        chart: {
            plotBackgroundColor: null,
            plotBorderWidth: null,
            plotShadow: false,
            renderTo: selector,
            type: 'pie'
        },
        title: {
            text: ''
        },
        tooltip: {
            formatter: function() {
                return this.key + ': <b>' + this.y + '</b>';
            },
            //pointFormat: '<b>{point.y}</b>',
            style: {
                fontSize: '15px'
            }
        },
        plotOptions: {
            pie: {
                allowPointSelect: true,
                cursor: 'pointer',
                dataLabels: {
                    enabled: true,
                    format: '<b>{point.name}</b>: {point.percentage:.1f} %',
                    style: {
                        color: (this.Highcharts.theme && this.Highcharts.theme.contrastTextColor) || 'black',
                        fontSize: '15px'
                    }
                }
            }
        },
        series: [{
            name: seriesName,
            colorByPoint: true,
            data: data
        }]
    }
  }
}

export default PieChart

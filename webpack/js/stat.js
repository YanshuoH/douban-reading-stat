import $ from 'jquery'
import api from './api'
import YearStat from './yearStat'
import MonthStat from './monthStat'
import PieChart from './pieChart'
import PosterWall from './posterWall'


class Stat {
  constructor() {
    this.yearNavContainer = $('#yearNav')
  }

  init() {
    this.yearStat = new YearStat()
    this.monthStat = new MonthStat()
    this.pieChart = new PieChart()
    this.posterWall = new PosterWall()

    // Retrieve content from api
    let userId = $('#userId').attr('data-id')
    api.getStat(userId, response => {
      this.handleStat(response)
    }, jqXHR => {
      // TODO
    })
  }

  handleStat(response) {
    // First of all, sort year array for order
    let yearArr = []
    let count = 0
    for (let year in response['result']) {
      yearArr.push(year)
      count ++
    }

    yearArr.sort((a, b) => b - a)
    yearArr = yearArr.slice(0, 7)
    for (let i = 0; i < yearArr.length; i ++) {
      this.yearNavContainer.append(`
        <li><a href="#${yearArr[i]}" id="${yearArr[i]}">${yearArr[i]}</a></li>
      `)
    }

    // Init by zero
    this.charts(response['result'], yearArr[0])
    $('#' + yearArr[0]).parent().addClass('active')

    // Listeners for all years
    for (let i = 0; i < yearArr.length; i ++) {
      let elem = $('#' + yearArr[i])
      elem.on('click', (e) => {
        e.preventDefault()
        this.charts(response['result'], yearArr[i])

        $('#yearNav > li').removeClass('active')
        elem.parent().addClass('active')
      })
    }
  }

  charts(result, year) {
    let havePreviousYear = result[year - 1] != undefined ? true : false
    this.yearStat.render(
      year, result[year]['count'],
      havePreviousYear == true ? result[year - 1]['count'] : 0
    )

    this.monthStat.render(result[year], result[year - 1])

    this.pieChart.renderRating(result[year]['rating'])
    this.pieChart.renderAuthor(result[year]['author'])
    this.pieChart.renderTranslator(result[year]['translator'])
    this.pieChart.renderPublisher(result[year]['publisher'])

    this.posterWall.bindClick(result[year]['posters'])
  }
}


export default Stat

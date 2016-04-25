import $ from 'jquery'

class YearStat {
  constructor() {
    this.yearCountYear = $('#yearCountYear')
    this.yearCountCount = $('#yearCountCount')
    this.yearCountPrevious = $('#yearCountPrevious')
  }

  render(year, totalCount, lastYearCount) {
    this.yearCountYear.html(year)
    this.yearCountCount.html(totalCount)
    this.yearCountPrevious.html(lastYearCount)
  }
}

export default YearStat

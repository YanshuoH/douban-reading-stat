import $ from 'jquery'

class PosterWall {
  constructor() {
    this.posterWallContainer = $('#posterWall')
    this.posterWallButton = $('#posterWallStart')
  }

  // Represent dataset[2015]['result']['posters']
  render(posters) {
    var container = this.posterWallContainer

    for (var i = 0; i < posters.length; i ++) {
      var poster = posters[i];
      container.append(`<img class="book-cover" src="${poster['link']}" alt="${poster['title']}">`)
    }
  }

  bindClick(posters) {
    this.posterWallContainer.empty()
    this.posterWallButton.unbind('click')

    let that = this
    this.posterWallButton.on('click', function() {
      that.render(posters)
    })
  }
}

export default PosterWall

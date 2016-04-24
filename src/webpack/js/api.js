import $ from 'jquery'

class Api {
  static getUser(url, doneFunc, failFunc) {
    let request = $.ajax({
      url: '/api/user?url=' + encodeURIComponent(url),
      contentType: 'application/json; charset=utf-8',
    }).done(doneFunc).fail(failFunc)
  }
}

export default Api

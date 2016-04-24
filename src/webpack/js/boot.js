import $ from 'jquery'

let url = 'https://www.douban.com/people/2274326/'
let request = $.ajax({
  url: '/api/user?url=' + encodeURIComponent(url),
  contentType: 'application/json; charset=utf-8',
}).done(res => {
  console.log(res)
}).fail((jqXHR, s) => {
  console.log(jqXHR)
  switch (jqXHR.status) {
    case 400:
      console.log(400)
      break
    default:
      console.log(404);
  }
});

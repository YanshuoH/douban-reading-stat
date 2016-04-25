import $ from 'jquery'
import form from './form'
import Stat from './stat'

$(() => {
  $(window).scroll(function () {
      let top = $(document).scrollTop()
      if(top > 50)
        $('#home > .navbar').removeClass('navbar-transparent')
      else
        $('#home > .navbar').addClass('navbar-transparent')
  });

  form()
  if ($('#userId').length > 0) {
    let stat = new Stat()
    stat.init()
  }
})

import $ from 'jquery'
import form from './form'

$(() => {
  $(window).scroll(function () {
      let top = $(document).scrollTop()
      if(top > 50)
        $('#home > .navbar').removeClass('navbar-transparent')
      else
        $('#home > .navbar').addClass('navbar-transparent')
  });

  form()
})

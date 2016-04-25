import $ from 'jquery'
import api from './api'

function form() {
  if (!$("#submit")) {
    return
  }

  let messageMapping = {
    '0': 'This web app will generate your personal reading statistics.',
    '1': 'Loading... <i class="fa fa-spinner fa-spin"></i> This may take a moment...',
    '404': 'Error: Can\'t find your profil using the given url',
    '400': 'Error: We may reach Douban API\'s rate limit (150 Request per day)'
  }

  let messageContainer = $("#formMessage")

  // Listener
  $("#submit").click(function() {
    let url = $('#profilUrlInput').val()
    messageContainer.html(messageMapping['1'])

    api.getUser(url, res => {
      document.location = res.next
    }, jqXHR => {
      switch (jqXHR.status) {
        case 400:
          messageContainer.html(messageMapping[jqXHR.status])
          break
        default:
          messageContainer.html(messageMapping[jqXHR.status])
      }
    });
  })
}

export default form

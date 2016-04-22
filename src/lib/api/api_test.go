package api

import (
  "fmt"
  "testing"
  "github.com/stretchr/testify/assert"

  "lib/mock"
)

func TestGetUser(t *testing.T) {
  assert := assert.New(t)

  username := "74783952"
  url := fmt.Sprintf(USER_URL, username)
  mock.PrepareResponseMock(url, mock.USER_RESPONSE_ASSET)

  _, status, err := GetUser(username)
  assert.Equal(200, status, "Existing user", "Should return a http 200 status")
  assert.Nil(err, "Existing user", "Should not occur an error")

  username = "3456789567895678956789"
  url = fmt.Sprintf(USER_URL, username)

  mock.PrepareResponseMock(url, mock.USER_404_RESPONSE_ASSET)

  _, status, err = GetUser("3456789567895678956789")
  assert.Equal(404, status, "Non-existing user", "Should return a http 404 status")
  assert.Nil(err, "Non-existing user", "Should not occur an error")
}

package api

import (
  "fmt"
  "testing"
  "github.com/stretchr/testify/assert"

  "github.com/YanshuoH/douban-reading-stat/lib/mock"
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

func TestGetUserBooks(t *testing.T) {
  assert := assert.New(t)

  username := "74783952"
  url1 := fmt.Sprintf(BOOK_COLLECTION_URL, username, 0)
  url2 := fmt.Sprintf(BOOK_COLLECTION_URL, username, 100)
  url3 := fmt.Sprintf(BOOK_COLLECTION_URL, username, 200)
  mock.PrepareResponseMock(url1, mock.BOOK_COLLECTION_RESPONSE_ASSET)
  mock.PrepareResponseMock(url2, mock.BOOK_COLLECTION_RESPONSE_ASSET_PAGE_2)
  mock.PrepareResponseMock(url3, mock.BOOK_COLLECTION_RESPONSE_ASSET_PAGE_3)

  _, status, err := GetUserBooks(username)
  assert.Equal(200, status, "Existing user for books", "Should return a http 200 status")
  assert.Nil(err, "Existing user for books", "Should not occur an error")

  username = "3456789567895678956789"
  url := fmt.Sprintf(BOOK_COLLECTION_URL, username, 0)

  mock.PrepareResponseMock(url, mock.BOOK_COLLECTION_404_RESPONSE_ASSET)

  _, status, err = GetUserBooks("3456789567895678956789")
  assert.Equal(404, status, "Non-existing user for books", "Should return a http 404 status")
  assert.Nil(err, "Non-existing user for books", "Should not occur an error")
}

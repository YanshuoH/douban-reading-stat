package models

import (
  "fmt"
  "lib/api"
  "lib/mock"
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestCastApiToModel(t *testing.T) {
  assert := assert.New(t)

  username := "74783952"
  url := fmt.Sprintf(api.USER_URL, username)
  mock.PrepareResponseMock(url, mock.USER_RESPONSE_ASSET)

  userData, _, _ := api.GetUser(username)

  url1 := fmt.Sprintf(api.BOOK_COLLECTION_URL, username, 0)
  url2 := fmt.Sprintf(api.BOOK_COLLECTION_URL, username, 100)
  url3 := fmt.Sprintf(api.BOOK_COLLECTION_URL, username, 200)
  mock.PrepareResponseMock(url1, mock.BOOK_COLLECTION_RESPONSE_ASSET)
  mock.PrepareResponseMock(url2, mock.BOOK_COLLECTION_RESPONSE_ASSET_PAGE_2)
  mock.PrepareResponseMock(url3, mock.BOOK_COLLECTION_RESPONSE_ASSET_PAGE_3)

  userBookData, _, _ := api.GetUserBooks(username)

  user := CastApiToModel(userData, userBookData)

  var aString string
  var anArray []interface{}
  assert.IsType(aString, user.UserId, "User Model", "Should be a string type for userId")
  assert.IsType(aString, user.Name, "User Model", "Should be a string type for name")
  assert.IsType(aString, user.Avatar, "User Model", "Should be a string type for url")
  assert.IsType(anArray, user.Books, "User Model", "Should be a array collection for books")
  assert.True(len(user.Books) > 0, "User Model", "Should have multiiple books")
}

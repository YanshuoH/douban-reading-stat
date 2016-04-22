package fetcher

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

const (
  URL_TEST_1 = "https://www.douban.com/people/74783952/"
  URL_TEST_2 = "https://www.douban.com/people/test/"
  URL_TEST_3 = "https://www.douban.com/people/test123/"
  URL_TEST_4 = "https://www.douban.com/people/"
  URL_TEST_5 = "https://www.douban.com/nein/74783952/"
)

func TestFetchUsername(t *testing.T) {
  assert := assert.New(t)

  username, err := FetchUsername(URL_TEST_1)
  assert.Nil(err, "Fetcher", "Should not be an error for int username")
  assert.Equal("74783952", username, "Fetcher", "Should be the user 74783952")

  username, err = FetchUsername(URL_TEST_2)
  assert.Nil(err, "Fetcher", "Should not be an error for string username")
  assert.Equal("test", username, "Fetcher", "Should be the user test")

  username, err = FetchUsername(URL_TEST_3)
  assert.Nil(err, "Fetcher", "Should not be an error for mix username")
  assert.Equal("test123", username, "Fetcher", "Should be the user test123")

  _, err = FetchUsername(URL_TEST_4)
  assert.NotNil(err, "Should be an error")
  assert.EqualError(err, "not found", "Fetcher", "Should be a not found error")

  _, err = FetchUsername(URL_TEST_5)
  assert.NotNil(err, "Should be an error")
  assert.EqualError(err, "not found", "Fetcher", "Should be a not found error")
}

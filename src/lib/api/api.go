package api

import (
  "fmt"
  "encoding/json"
  "net/http"
  "io/ioutil"
)

const (
  USER_URL = "https://api.douban.com/v2/user/%s"
  BOOK_COLLECTION_URL = "https://api.douban.com/v2/book/user/%s/collections?count=100&start=%d"
)

func GetUser(username string) (map[string]interface{}, int, error) {
  // Make url
  url := fmt.Sprintf(USER_URL, username)

  // Requesting API
  response, err := http.Get(url)

  return responseHandler(response, err)
}

func GetUserBooks(username string) ([]interface{}, int, error) {
  var result []interface{}
  var data map[string]interface{}
  var status int
  var err error

  start := 0
  for {
      // Make url
      url := fmt.Sprintf(BOOK_COLLECTION_URL, username, start)

      var response *http.Response
      // Requesting API
      response, err = http.Get(url)

      data, status, err = responseHandler(response, err)

      if status != 200 || err != nil {
        break
      }

      collections := data["collections"].([]interface{})
      if len(collections) == 0 {
        break
      }

      result = append(result, collections...)

      if (data["start"].(float64) + data["count"].(float64)) < data["total"].(float64) {
        start += 100
      } else {
        break
      }
  }

  return result, status, err
}

func responseHandler(response *http.Response, err error) (map[string]interface{}, int, error) {
  // Holder data var
  var data map[string]interface{}
  var status int = 0

  // General error handling
  if err != nil {
    return data, status, err
  }

  // Read status
  status = response.StatusCode

  // Read body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return data, status, err
	}

  byt := []byte(body)
  err = json.Unmarshal(byt, &data)

  return data, status, err
}

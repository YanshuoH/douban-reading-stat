package statistic

import (
  "fmt"
  "io/ioutil"
  "testing"
  "github.com/stretchr/testify/assert"

  "github.com/YanshuoH/douban-reading-stat/models"
  "github.com/YanshuoH/douban-reading-stat/lib/mock"
  "github.com/YanshuoH/douban-reading-stat/lib/api"
)

const (
  USERNAME = "2274326"
  USER_ASSET = `{"loc_id":"108288","name":"优","created":"2008-03-02 11:04:55","is_banned":false,"is_suicide":false,"loc_name":"北京","avatar":"https://img3.doubanio.com\/icon\/u2274326-10.jpg","signature":"一个被低级趣味拖住死缠烂打的人","uid":"2274326","alt":"http:\/\/www.douban.com\/people\/2274326\/","desc":"丑的真诚 （东北人研究学会）\n\n不谈政治了，门槛太低。","type":"user","id":"2274326","large_avatar":"https://img3.doubanio.com\/icon\/up2274326-10.jpg"}`
)

var Responses []string

// For the sake of my IDE, read file content instead of putting all response in const
func setup() {
  baseDir := "../mock/dataset"
  files, err := ioutil.ReadDir(baseDir)
  checkErr(err)

  dataArray := make([]string, len(files))

  for i, f := range files {
    relativePath := baseDir + "/" + f.Name()
    dat, err := ioutil.ReadFile(relativePath)
    checkErr(err)

    dataArray[i] = string(dat)
  }

  // Prepare books response mock
  for i, content := range dataArray {
    url := fmt.Sprintf(api.BOOK_COLLECTION_URL, USERNAME, (i * 100))
    mock.PrepareResponseMock(url, content)
  }

  // Prepare user response mock
  url := fmt.Sprintf(api.USER_URL, USERNAME)
  mock.PrepareResponseMock(url, USER_ASSET)
}

func TestStat(t *testing.T)  {
  setup()
  assert := assert.New(t)
  userData, _, _ := api.GetUser(USERNAME)
  userBookData, _, _ := api.GetUserBooks(USERNAME)
  user := models.CastApiToModel(userData, userBookData)

  // regroup by year
  result := Stat(user)
  assert.Equal(9, len(result), "Statistic", "Should divide to 9 groups")
  assert.Equal(4, result["2008"].Count, "Statistic", "Should have read 4 books in 2008")
  expectedRating := map[string]int{
    "1": 0,
    "2": 0,
    "3": 1,
    "4": 1,
    "5": 1,
  }
  assert.Equal(expectedRating, result["2008"].Rating, "Statistic", "Should have above rating statistic")
}

package controllers

import (
  "fmt"
  "time"
  "strconv"
  "github.com/gin-gonic/gin"
  "gopkg.in/mgo.v2"

  "models"
  userRepository "repository/user"
  "lib/fetcher"
  "lib/api"
  "lib/statistic"
)

const (
  RECYCLE_PERIOD = 24
)

func html404(c *gin.Context) {
  c.HTML(404, "404", gin.H{
    "errors": "No man's land",
  })
}

func retrieveUser(c *gin.Context, username string) (int, *models.User, error, string) {
  found := false

  // Connection
  con := c.MustGet("db").(*mgo.Database)

  // Retrieve user in DB
  userModel, err := userRepository.Get(username, con)

  if err == nil && len(userModel.UserId) > 0 {
    found = true
    now := time.Now()
    diff := now.Sub(userModel.UpdatedOn)
    // If in recycle period
    if diff.Hours() < RECYCLE_PERIOD {
      return 200, userModel, nil, "ok"
    }
  }

  // Get user
  userInfo, status, err := api.GetUser(username)
  if err != nil || status != 200 {
    return status, &models.User{}, err, "Douban user API return a status of " + strconv.Itoa(status)
  }

  // Get collections
  collections, status, err := api.GetUserBooks(username)
  if err != nil || status != 200 {
    return status, &models.User{}, err, "Douban book collections API return a status of " + strconv.Itoa(status)
  }

  var stat map[string]*models.StatEntity

  // Cast to user model
  newUserModel := models.CastApiToModel(userInfo, collections)

  stat = statistic.Stat(newUserModel)
  newUserModel.Stat = stat

  if !found {
    _, err = userRepository.Save(*newUserModel, con)
  } else {
    _, err = userRepository.Update(*newUserModel, con)
  }

  return 200, newUserModel, nil, ""
}


/**
 * Public
 */
func Index(c *gin.Context) {
  c.HTML(200, "index.html", gin.H{})
}

func StatPage(c *gin.Context) {
  userId := c.Param("userId")
  if userId == "" {
    html404(c)
    return
  }

  status, userModel, err, _ := retrieveUser(c, userId)
  if err != nil || status != 200 {
    html404(c)
    return
  }

  c.HTML(200, "stat.html", gin.H{
    "status": 200,
    "userId": userModel.UserId,
  })
}

/**
 * API
 */
func SearchUser(c *gin.Context) {
  url := c.Query("url")
  if url == "" {
    c.JSON(404, gin.H{
      "status": 404,
      "error": "invalid url",
    })
    return
  }

  username, err := fetcher.FetchUsername(url)

  if err != nil {
    c.JSON(404, gin.H{
      "status": 404,
      "msg": "invalid url",
    })
    return
  }

  status, userModel, err, msg := retrieveUser(c, username)
  if err != nil || status != 200 {
    c.JSON(status, gin.H{
      "status": status,
      "msg": msg,
    })
    return
  }

  c.JSON(200, gin.H{
    "status": 200,
    "next": "/user/" + userModel.UserId,
  })
}

func GetUser(c *gin.Context) {
  userId := c.Param("userId")
  fmt.Println(userId)
  if userId == "" {
    c.JSON(404, gin.H{
      "status": 404,
      "error": "invalid url",
    })
    return
  }

  status, userModel, err, msg := retrieveUser(c, userId)
  if err != nil || status != 200 {
    c.JSON(status, gin.H{
      "status": status,
      "msg": msg,
    })
    return
  }

  res := make(map[string]interface{})
  for key, value := range userModel.Stat {
    res[key] = value
  }

  c.JSON(200, gin.H{
    "userId": userModel.UserId,
    "username": userModel.Name,
    "avatar": userModel.Avatar,
    "result": res,
  })
}

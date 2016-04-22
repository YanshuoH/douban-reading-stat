package controllers

import (
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

func Index()  {

}

func Entry(c *gin.Context)  {
  // For testing
  url := "https://www.douban.com/people/2274326/"
  found := false
  username, err := fetcher.FetchUsername(url)
  if err != nil {
    c.JSON(500, gin.H{
      "status": 500,
      "msg": "Cannot recognize username/id using given url/string",
    })
    return
  }

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
      res := make(map[string]interface{})
      for key, value := range userModel.Stat {
        res[key] = value
      }

      c.JSON(200, gin.H{
        "result": res,
      })
      return
    }
  }

  // Get user
  userInfo, status, err := api.GetUser(username)
  if err != nil || status != 200 {
    c.JSON(status, gin.H{
      "status": status,
      "msg": "Douban user API return a status of " + strconv.Itoa(status),
      "error": err,
    })
    return
  }

  // Get collections
  collections, status, err := api.GetUserBooks(username)
  if err != nil || status != 200 {
    c.JSON(status, gin.H{
      "status": status,
      "msg": "Douban book collections API return a status of " + strconv.Itoa(status),
      "error": err,
    })
    return
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

  res := make(map[string]interface{})
  for key, value := range stat {
    res[key] = value
  }

  c.JSON(200, gin.H{
    "result": res,
  })

}

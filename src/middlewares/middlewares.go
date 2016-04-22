package middlewares

import (
  "github.com/gin-gonic/gin"
  "db"
)

func ConnectDb(c *gin.Context)  {
  s := db.Session.Clone()

  defer s.Close()

  c.Set("db", s.DB(db.Mongo.Database))
  c.Next()
}

func ErrorHandler(c *gin.Context)  {
  c.Next()
  /*
   * TODO
   */
}

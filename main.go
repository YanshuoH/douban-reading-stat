package main

import (
  // "net/http"
  "os"

  "github.com/gin-gonic/gin"

  "github.com/YanshuoH/douban-reading-stat/db"
  "github.com/YanshuoH/douban-reading-stat/controllers"
  "github.com/YanshuoH/douban-reading-stat/middlewares"
)

const (
	// Port at which the server starts listening
	Port = "3000"
)

func init() {
	db.Connect()
}

func main() {
  // Configure Gin
  router := gin.Default()

  // Set html render options
  router.LoadHTMLGlob("./templates/*")

  router.RedirectTrailingSlash = true
	router.RedirectFixedPath = true

  // Middlewares
  router.Use(middlewares.ConnectDb)

  // Statics
  router.Static("/public", "./public")

  // Routing list
  router.GET("/", controllers.Index)
  router.GET("/user/:userId", controllers.StatPage)
  router.GET("/api/user", controllers.SearchUser)
  router.GET("/api/user/:userId", controllers.GetUser)

  // Start listening
  port := Port
  if len(os.Getenv("PORT")) > 0 {
    port = os.Getenv("PORT")
  }
  router.Run(":" + port)
}

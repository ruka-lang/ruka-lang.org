package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/**/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "warp.index.tmpl.html", nil)
	})
  
    router.GET("/drive", func(c *gin.Context) {
		c.HTML(http.StatusOK, "drive.index.tmpl.html", nil)
	})
	
    router.GET("/whdl", func(c *gin.Context) {
		c.HTML(http.StatusOK, "whdl.index.tmpl.html", nil)
	})
  
    router.GET("/about", func(c *gin.Context) {
		c.HTML(http.StatusOK, "about.index.tmpl.html", nil)
	})

	router.Run(":" + port)
}

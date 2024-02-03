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
		c.HTML(http.StatusOK, "amulet.index.tmpl.html", nil)
	})
  
    router.GET("/ritual", func(c *gin.Context) {
		c.HTML(http.StatusOK, "ritual.index.tmpl.html", nil)
	})
	
    router.GET("/ahdl", func(c *gin.Context) {
		c.HTML(http.StatusOK, "ahdl.index.tmpl.html", nil)
	})
  
    router.GET("/about", func(c *gin.Context) {
		c.HTML(http.StatusOK, "about.index.tmpl.html", nil)
	})

	router.Run(":" + port)
}

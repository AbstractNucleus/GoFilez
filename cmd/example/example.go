package main

import (
	"net/http"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

func multiRender() multitemplate.Render {
	r := multitemplate.New()
	r.AddFromFiles("index", "index.html", "index.html")
	return r
}

func main() {
	router := gin.Default()
	router.HTMLRender = multiRender()

	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index", gin.H{
			"title": "Html5 Template Engine",
		})
	})

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.Run()
}

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("pie.html")
	r.GET("/images", images)

	r.Run(":9000")
}

func images(c *gin.Context) {
	c.HTML(http.StatusOK, "pie.html", gin.H{})
}

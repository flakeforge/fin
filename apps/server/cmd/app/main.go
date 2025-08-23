package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	println("Hello, Fin server!")

	r := gin.Default()
  r.GET("/", func (c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "Hello, Fin server!",
    })
  })
  r.Run() // listen and serve on
}

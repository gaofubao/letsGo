package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/header", HandleHeader)
	r.GET("/ping", HandlePing)

	r.Run()
}

func HandlePing(c *gin.Context)  {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func HandleHeader(c *gin.Context) {
	headerJson, _ := json.Marshal(c.Request.Header)
	c.Writer.WriteString(string(headerJson))
}

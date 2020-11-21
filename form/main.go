package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// curl -i -X POST 127.0.0.1:8080/users \
// -F "message=hello"
func main() {
	r := gin.Default()
	r.POST("/users", func(c *gin.Context) {
		message := c.PostForm("message")
		c.JSON(http.StatusOK, gin.H{
			"message": message,
		})
	})
	r.Run()
}

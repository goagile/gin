package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// curl -i -X POST 127.0.0.1:8080/users/ivan
// curl -i -X POST 127.0.0.1:8080/users/ivan/hello
func main() {
	r := gin.Default()
	r.GET("/users/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "My name is "+name)
	})
	r.GET("/users/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		c.String(http.StatusOK, "My name is "+name+" action "+action)
	})
	r.Run()
}

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// curl -i \
// -X GET 127.0.0.1:8080/users?page=1&sort=-name
func main() {
	r := gin.Default()
	r.GET("/users", func(c *gin.Context) {
		page := c.DefaultQuery("page", "0")
		sort := c.Query("sort")
		c.String(http.StatusOK, "sort: %v page: %v\n", sort, page)
	})
	r.Run()
}

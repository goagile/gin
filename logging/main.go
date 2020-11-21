package main

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.DisableConsoleColor()
	f, _ := os.Create("log.log")
	defer f.Close()
	gin.DefaultWriter = io.MultiWriter(f)
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "HELLO")
	})
	r.Run(":8080")
}

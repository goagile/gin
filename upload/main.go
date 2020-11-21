package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// curl -i \
// -X POST 127.0.0.1/8080/upload \
// -F "file=@x.file"
func main() {
	r := gin.Default()
	r.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		c.JSON(http.StatusOK, gin.H{
			"filename": file.Filename,
			"header":   file.Header,
			"size":     file.Size,
		})
	})
	r.Run()
}

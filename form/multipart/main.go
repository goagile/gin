package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// curl -i -X POST 127.0.0.1:8080/upload \
// -F "upload[]=@1.txt"
// -F "upload[]=@2.txt"
// -H "Content-Type: multipart/form-data"
func main() {
	r := gin.Default()
	r.POST("/upload", func(c *gin.Context) {
		form, _ := c.MultipartForm()
		files := form.File["upload[]"]
		for _, file := range files {
			log.Println(file.Filename)
		}
		c.JSON(http.StatusOK, gin.H{
			"count": len(files),
		})
	})
	r.Run()
}

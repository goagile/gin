package main

import (
	"github.com/gin-gonic/gin"
	"github.com/goagile/gin/books"
)

func main() {
	r := gin.Default()

	r.POST("/books", books.Create)       // CREATE	(C)
	r.GET("/books/:id", books.Find)      // READ	(R)
	r.PUT("/books/:id", books.Update)    // UPDATE	(U)
	r.DELETE("/books/:id", books.Delete) // DELETE	(D)
	r.GET("/books", books.FindAll)       // LIST	(L)

	r.Run(":8080")
}

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/goagile/gin/controller"
)

func main() {
	r := gin.Default()

	r.POST("/books", controller.CreateBook)    // CREATE	(C)
	r.GET("/books/:id", controller.FindBook)   // READ		(R)
	r.PUT("/books/:id", controller.UpdateBook) // UPDATE	(U)
	// r.DELETE("/books/:id", deleteBook) // DELETE	(D)
	r.GET("/books", controller.FindBooks) // LIST		(L)

	r.Run(":8080")
}

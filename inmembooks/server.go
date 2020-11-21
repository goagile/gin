package main

import (
	"github.com/gin-gonic/gin"
	"github.com/goagile/gin/inmembooks/books"
)

func init() {
	gin.SetMode(gin.DebugMode)
}

func main() {
	setupServer().Run()
}

func setupServer() *gin.Engine {
	r := gin.Default()

	// (C) CREATE
	// POST host:port/books
	r.POST("/books", books.Create)

	// (R) READ
	// GET host:port/books/32
	r.GET("/books/:id", books.Find)

	// (U) UPDATE
	// PUT host:port/books/32
	r.PUT("/books/:id", books.Update)

	// (D) DELETE
	// DELETE host:port/books/32
	r.DELETE("/books/:id", books.Delete)

	// (L) LIST
	// GET host:port/books?perpage=10&page=2
	r.GET("/books", books.FindAll)

	return r
}

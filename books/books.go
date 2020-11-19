package books

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/goagile/gin/books/book"
	"github.com/goagile/gin/books/db"
)

// Create - create new book
func Create(c *gin.Context) {
	var r CreateRequest
	if err := c.BindJSON(&r); err != nil {
		log.Println("CreateBook BindJSON", err)
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": "fail to create book"},
		)
		return
	}
	b := newBook(r)
	db.Save(b)
	c.JSON(http.StatusCreated, gin.H{"data": b})
}

// CreateRequest - request to create book
type CreateRequest struct {
	Author string `json:"author"`
	Title  string `json:"title"`
}

// newBook - return new Book
func newBook(r CreateRequest) *book.Book {
	b := new(book.Book)
	b.ID = db.NextID()
	b.Author = r.Author
	b.Title = r.Title
	return b
}

// Find one book by ID
func Find(c *gin.Context) {
	id, err := book.IDFromString(c.Param("id"))
	if err != nil {
		log.Println("FindBook IDFromString", err)
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": fmt.Sprintf("not found book by id %q", id)},
		)
		return
	}
	b := db.Find(id)
	c.JSON(http.StatusOK, gin.H{"data": b})
}

// Update - update existing book
func Update(c *gin.Context) {
	var r UpdateRequest
	c.BindJSON(&r)
	id, err := book.IDFromString(c.Param("id"))
	if err != nil {
		log.Println("UpdateBook IDFromString", err)
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": fmt.Sprintf("not found book to update by id %q", id)},
		)
		return
	}
	b := db.Find(id)
	b = updatedBook(r, b)
	db.Save(b)
	c.JSON(http.StatusOK, gin.H{"data": b})
}

// updatedBook - update book from request
func updatedBook(r UpdateRequest, b *book.Book) *book.Book {
	if "" != r.Author {
		b.Author = r.Author
	}
	if "" != r.Title {
		b.Title = r.Title
	}
	return b
}

// UpdateRequest - request to update book
type UpdateRequest struct {
	Author string `json:"author"`
	Title  string `json:"title"`
}

// FindAll - return all books
func FindAll(c *gin.Context) {
	bs := db.FindAll()
	c.JSON(http.StatusOK, gin.H{"data": bs})
}

// Delete - delete book by ID
func Delete(c *gin.Context) {
	id, err := book.IDFromString(c.Param("id"))
	if err != nil {
		log.Println("DeleteBook IDFromString", err)
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": fmt.Sprintf("not found book to delete by id %q", id)},
		)
		return
	}
	db.Delete(id)
	c.JSON(http.StatusOK, gin.H{"data": id})
}

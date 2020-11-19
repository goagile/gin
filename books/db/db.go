package db

import (
	"sort"
	"sync"

	"github.com/goagile/gin/books/book"
)

var (
	// save id
	idMu sync.Mutex
	id   int64

	// save db
	dbMu sync.Mutex
	db   = make(map[int64]*book.Book)
)

// Save - method save book to storage
func Save(b *book.Book) int64 {
	dbMu.Lock()
	db[b.ID] = b
	dbMu.Unlock()
	return b.ID
}

// FindAll - return all books
func FindAll() []*book.Book {
	bs := make([]*book.Book, 0)
	dbMu.Lock()
	for _, b := range db {
		bs = append(bs, b)
	}
	sort.Sort(book.ByID(bs))
	dbMu.Unlock()
	return bs
}

// FindPerPage - find all books skip and limit result
func FindPerPage(perpage, page int) []*book.Book {
	bs := FindAll()
	return bs
}

// Find - find book by ID
func Find(id int64) *book.Book {
	return db[id]
}

// Delete - method delete book from storage
func Delete(id int64) {
	dbMu.Lock()
	delete(db, id)
	dbMu.Unlock()
}

// NextID - returns next books ID
func NextID() int64 {
	incID()
	return getID()
}

func incID() {
	idMu.Lock()
	id++
	idMu.Unlock()
}

func getID() int64 {
	var v int64
	idMu.Lock()
	v = id
	idMu.Unlock()
	return v
}

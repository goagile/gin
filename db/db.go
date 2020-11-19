package db

import (
	"sync"

	"github.com/goagile/gin/book"
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
	dbMu.Unlock()
	return bs
}

// Find - find book by ID
func Find(id int64) *book.Book {
	return db[id]
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

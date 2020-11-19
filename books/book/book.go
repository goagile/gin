package book

import "strconv"

// New - book constructor
func New(id int64, title, author string) *Book {
	b := new(Book)
	b.Title = title
	b.Author = author
	return b
}

// Book - model entity
type Book struct {
	ID     int64  `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

// ByID books sorter
type ByID []*Book

func (s ByID) Less(i, j int) bool {
	return (s[i].Title < s[j].Title)
}

func (s ByID) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByID) Len() int {
	return len(s)
}

// IDFromString - create book ID from string
func IDFromString(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}

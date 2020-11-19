package db

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/goagile/gin/books/book"
)

var bs = makeBooks()

func makeBooks() []*book.Book {
	bs := []*book.Book{}
	for i, r := range "ABCDE" {
		n := string(r)
		b := &book.Book{
			ID:     NextID(),
			Author: n,
			Title:  fmt.Sprintf("%v_%v", n, i),
		}
		bs = append(bs, b)
	}
	return bs
}

func saveBooks() {
	for _, b := range bs {
		Save(b)
	}
}

func Test_FindAll(t *testing.T) {
	saveBooks()
	want := makeBooks()

	got := FindAll()

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

package db

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/goagile/gin/inmembooks/books/book"
)

var bs = makeBooks("ABC")

func makeBooks(s string) []*book.Book {
	bs := []*book.Book{}
	for i, r := range s {
		b := makeBook(i, string(r))
		bs = append(bs, b)
	}
	return bs
}

func makeBook(i int, name string) *book.Book {
	b := new(book.Book)
	b.ID = NextID()
	b.Author = name
	b.Title = fmt.Sprintf("%v_%v", name, i+1)
	return b
}

func saveBooks() {
	for _, b := range bs {
		Save(b)
	}
}

func TestMain(m *testing.M) {
	saveBooks()
	code := m.Run()
	Clear()
	os.Exit(code)
}

//
// FindAll
//
func Test_FindAll(t *testing.T) {
	want := bs

	got := FindAll()

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

//
// FindPerPage
//
func Test_FindPerPage_Zero(t *testing.T) {
	want := []*book.Book{}

	got := FindPerPage(0, 0)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

func Test_FindPerPage_LtZero(t *testing.T) {
	want := []*book.Book{}

	got := FindPerPage(-1, -1)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

func Test_FindPerPage_GtLen(t *testing.T) {
	want := []*book.Book{}

	got := FindPerPage(10, 10)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

func Test_FindPerPage_1_1(t *testing.T) {
	want := []*book.Book{
		{ID: int64(1), Author: "A", Title: "A_1"},
	}

	got := FindPerPage(1, 1)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

func Test_FindPerPage_1_2(t *testing.T) {
	want := []*book.Book{
		{ID: int64(2), Author: "B", Title: "B_2"},
	}

	got := FindPerPage(1, 2)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

func Test_FindPerPage_1_3(t *testing.T) {
	want := []*book.Book{
		{ID: int64(3), Author: "C", Title: "C_3"},
	}

	got := FindPerPage(1, 3)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

func Test_FindPerPage_1_4(t *testing.T) {
	want := []*book.Book{}

	got := FindPerPage(1, 4)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

func Test_FindPerPage_2_1(t *testing.T) {
	want := []*book.Book{
		{ID: int64(1), Author: "A", Title: "A_1"},
		{ID: int64(2), Author: "B", Title: "B_2"},
	}

	got := FindPerPage(2, 1)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

func Test_FindPerPage_2_2(t *testing.T) {
	want := []*book.Book{
		{ID: int64(3), Author: "C", Title: "C_3"},
	}

	got := FindPerPage(2, 2)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

func Test_FindPerPage_2_3(t *testing.T) {
	want := []*book.Book{}

	got := FindPerPage(2, 3)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

func Test_FindPerPage_3_1(t *testing.T) {
	want := bs

	got := FindPerPage(3, 1)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

func Test_FindPerPage_3_2(t *testing.T) {
	want := []*book.Book{}

	got := FindPerPage(3, 2)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

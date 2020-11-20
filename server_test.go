package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/goagile/gin/books/book"
	"github.com/goagile/gin/books/db"
)

const ajson = "application/json"

func init() {
	gin.SetMode(gin.TestMode)
}

var server *httptest.Server

func TestMain(m *testing.M) {
	server = httptest.NewServer(setupServer())
	code := m.Run()
	server.Close()
	os.Exit(code)
}

func uri(path string) string {
	return fmt.Sprintf("%v%v", server.URL, path)
}

var warnpeace = &book.Book{
	ID:     1,
	Author: "Leo Tolstoy",
	Title:  "War and Peace",
}

// POST
func Test_PostBooks(t *testing.T) {
	want := &postRequest{warnpeace}

	r, err := http.Post(uri("/books"), ajson, dump(want))
	if err != nil {
		t.Fatal("Post", err)
	}
	defer r.Body.Close()

	got := fromBody(r)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("\nwant:%v\ngot:%v", want, got)
	}

	db.Clear()
}

type postRequest struct {
	Data *book.Book `json:"data"`
}

func dump(r *postRequest) *bytes.Reader {
	j, err := json.Marshal(r.Data)
	if err != nil {
		log.Fatal("dump Marshal", err)
	}
	return bytes.NewReader(j)
}

func fromBody(r *http.Response) *postRequest {
	var b *postRequest
	byt, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal("fromBody ReadAll", err)
	}
	if err := json.Unmarshal(byt, &b); err != nil {
		log.Fatal("Unmarshal", err)
	}
	return b
}

// LIST
func Test_ListBooks(t *testing.T) {
	want := map[string]interface{}{
		"data": []interface{}{},
	}

	r, err := http.Get(uri("/books"))
	if err != nil {
		t.Fatal("Get", err)
	}
	defer r.Body.Close()

	var got map[string]interface{}
	b, err := ioutil.ReadAll(r.Body)
	if err := json.Unmarshal(b, &got); err != nil {
		t.Fatal("Unmarshal", err)
	}

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("\nwant:%v\ngot:%v", want, got)
	}
}

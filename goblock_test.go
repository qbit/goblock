package goblock

import (
	"log"
	"testing"
	"fmt"
	"net/http"
	"net/http/httptest"
)

func TestGet(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hi")
	}))

	defer ts.Close()

	data, err := get(ts.URL)
	if err != nil {
		log.Fatal(err)
	}

	var answer string = string(data)

	if answer != "Hi" {
		log.Fatal("Not OK!")
	} else {
		log.Println("OK")
	}
}

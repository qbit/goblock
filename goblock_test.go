package goblock

import (
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
		fmt.Printf("%v\n", err)
		t.FailNow()
	}

	var answer = string(data)

	if answer != "Hi\n" {
		fmt.Println("Not OK!")
		t.FailNow()
	} else {
		fmt.Println("OK")
	}
}

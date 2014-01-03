package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"compress/gzip"
	"io/ioutil"
	"testing"
)

func fileContains(t *testing.T, src string, contents string) (bool) {
	buf, err := ioutil.ReadFile(src)
	if err != nil {
		fmt.Printf("%v\n", err)
		t.FailNow()
	}

	str := string(buf)

	if str == contents {
		return true
	} else {
		return false
	}
}

func TestGet(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var b bytes.Buffer

		gz := gzip.NewWriter(&b)
		gz.Write([]byte("Hi"))

		defer gz.Close()

		var s = string(b.Bytes())

		fmt.Fprintln(w, s)
	}))

	defer ts.Close()

	bytes, err := get(ts.URL, "/tmp/awesome")
	fmt.Println("put %d into /tmp/awesome", bytes)
	if err != nil {
		fmt.Printf("%v\n", err)
		t.FailNow()
	}

	if fileContains(t, "/tmp/awesome", "Hi\n") {
		fmt.Println("OK")
	} else {
		fmt.Println("Not OK!")
		t.FailNow()
	}
}

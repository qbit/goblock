package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func fileContains(t *testing.T, src string, contents string) bool {
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
		b := bytes.NewBufferString("Hi")
		gz := gzip.NewWriter(w)

		defer gz.Close()

		_, err := io.Copy(gz, b)
		if err != nil {
			fmt.Printf("Failed to write gzip response: %v\n", err)
		}
	}))

	defer ts.Close()

	os.Remove("/tmp/awesome")

	bytes, err := get(ts.URL, "/tmp/awesome")
	fmt.Printf("put %d into /tmp/awesome\n", bytes)
	if err != nil {
		fmt.Printf("%v\n", err)
		t.FailNow()
	}

	if fileContains(t, "/tmp/awesome", "Hi") {
		fmt.Println("OK")
	} else {
		fmt.Println("Not OK!")
		t.FailNow()
	}
}

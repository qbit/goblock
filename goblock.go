package main

import (
	"compress/gzip"
	"github.com/gokyle/goconfig"
	"io"
	"log"
	"net/http"
	"os"
)

func get(src string, dest string) (int64, error) {
	resp, err := http.Get(src)
	errr(err, "Can't make http request")

	file, err2 := os.Create(dest)
	errr(err2, "Can't create file!")

	defer resp.Body.Close()

	gz, err3 := gzip.NewReader(resp.Body)
	errr(err3, "Can't uncompress file!")

	defer gz.Close()
	defer file.Close()

	return io.Copy(file, gz)
}

func errr(e error, msg string) {
	if e != nil {
		log.Printf("[!]: %s - %s", msg, e)
	}
}

func main() {
	var conf, err = goconfig.ParseFile("config.ini")
	var url = conf["global"]["url"]
	var params string = "?"

	for key, val := range conf["params"] {
		params = params + key + "=" + val + "&"
	}

	errr(err, "Can't parse config file!")
	log.Printf("Getting lists from: %s", url)

	for key, val := range conf["list"] {
		var full_url = url + params + "list=" + val
		var file_name = "/tmp/" + key

		log.Printf("downloading %s", key)
		written, err2 := get(full_url, file_name)
		errr(err2, "Can't write file!")

		log.Printf("%d written to %s", written, file_name)
	}
}

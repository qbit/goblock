package main

import (
	"compress/gzip"
	"github.com/gokyle/goconfig"
	"io"
	"log"
	"net/http"
	"os"
)

var elog = log.New(os.Stderr, "[!]: ", 0)

func get(src, dest string) (int64, error) {
	resp, err := http.Get(src)
	if err != nil {
		elog.Printf("%v - %v", "Can't make http request", err)
	}

	file, err := os.OpenFile(dest, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		elog.Printf("%v - %v", "Can't create file!", err)
	}

	defer resp.Body.Close()

	gz, err := gzip.NewReader(resp.Body)
	if err != nil {
		elog.Printf("%v - %v", "Can't uncompress file!", err)
	}

	defer gz.Close()
	defer file.Close()

	// n, err := io.Copy(file, gz)
	// return n, err
	return io.Copy(file, gz)
}

func main() {
	conf, err := goconfig.ParseFile("config.ini")
	if err != nil {
		elog.Printf("%v - %v", "Can't parse config file!", err)
	}

	url := conf["global"]["url"]
	file_name := conf["global"]["destination"]
	params := "?"

	os.Remove(file_name)

	for key, val := range conf["params"] {
		params = params + key + "=" + val + "&"
	}

	log.Printf("Getting lists from: %s", url)

	for key, val := range conf["list"] {
		full_url := url + params + "list=" + val

		log.Printf("downloading %s", key)
		_, err := get(full_url, file_name)
		if err != nil {
			elog.Printf("%v - %v", "Can't write file!", err)
		}
	}
}

package main

import (
	"bufio"
	"github.com/gokyle/goconfig"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func get(src string) ([]byte, error) {
	resp, err := http.Get(src)
	errr(err, "Can't make http request")

	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func write(data []byte, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}

	defer file.Close()

	w := bufio.NewWriter(file)
	file.Write(data)

	return w.Flush()
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
		log.Printf("downloading %s", key)
		data, err := get(full_url)
		write(data, "/tmp/"+key+".gz")
		errr(err, "Error getting "+full_url)
	}
}

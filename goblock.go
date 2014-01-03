package goblock

import (
	"bufio"
	"github.com/gokyle/goconfig"
	"io/ioutil"
	"log"
	// "compress/gzip"
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

	// close this britch when we are done with parent func
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
		var file_name = "/tmp" + key + ".gz"

		// need to change this into a single piped gzip operation
		log.Printf("downloading %s", key)
		data, err := get(full_url)

		write(data, file_name)
		errr(err, "Error getting "+full_url)
	}
}

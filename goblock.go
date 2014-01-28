package main

import (
	"compress/gzip"
	"flag"
	"github.com/gokyle/goconfig"
	"io"
	"log"
	"net/http"
	"os"
)

func get(src, dest string) (int64, error) {
	resp, err := http.Get(src)
	if err != nil {
		log.Fatalf("Can't make http request - %v", err)
	}

	file, err := os.OpenFile(dest, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatalf("Can't create file! - %v", err)
	}

	defer resp.Body.Close()

	gz, err := gzip.NewReader(resp.Body)
	if err != nil {
		log.Fatalf("Can't uncompress file! - %v", err)
	}

	defer gz.Close()
	defer file.Close()

	return io.Copy(file, gz)
}

func main() {
	var cfile = flag.String("config", "config.ini", "Full path to config file.")
	var debug = flag.Bool("v", false, "show logging info.")
	flag.Parse()
	conf, err := goconfig.ParseFile(*cfile)
	if err != nil {
		log.Fatalf("Can't parse config file! - %v", err)
	}

	url := conf["global"]["url"]
	dfile_name := conf["global"]["destination"]
	params := "?"

	os.Remove(dfile_name)

	for key, val := range conf["params"] {
		params = params + key + "=" + val + "&"
	}

	if *debug {
		log.Printf("Getting lists from: %s", url)
	}

	for key, val := range conf["list"] {
		full_url := url + params + "list=" + val

		if *debug {
			log.Printf("downloading %s", key)
		}
		_, err := get(full_url, dfile_name)
		if err != nil {
			if *debug {
				log.Fatalf("Can't write file! - %v", err)
			}
		}
	}
}

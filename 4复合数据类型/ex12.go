package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type comic struct {
	Num        int
	Transcript string
}

func main() {
	var xkcd comic
	searchTerm := strings.Join(os.Args[1:], " ")
	if _, err := os.Stat("index/xkcd1.json"); err != nil {
		createComicIndex() // Create offline index of web comic xkcd.
	}
	for i := 1; ; i++ { // Loop until end of sequential files in index.
		file, err := os.Open(fmt.Sprintf("index/xkcd%d.json", i))
		if err != nil {
			if os.IsNotExist(err) {
				break // End of sequential files in offline index.
			}
			log.Fatal(err)
		}
		if err := json.NewDecoder(file).Decode(&xkcd); err != nil {
			file.Close()
			if i == 404 {
				continue // Skip 404 - Not Found
			}
			log.Fatal(err)
		}
		file.Close()
		if strings.Contains(strings.ToLower(xkcd.Transcript), strings.ToLower(searchTerm)) {
			fmt.Printf("\nFound search term %#v in ", searchTerm)
			fmt.Printf("https://xkcd.com/%d/\n", i)
			fmt.Println(xkcd.Transcript)
		}
	}
}

func createComicIndex() {
	fmt.Print("Creating xkcd offline index... ")
	var lastComic comic
	r, err := http.Get("https://xkcd.com/info.0.json") // Most recent comic.
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&lastComic); err != nil {
		log.Fatal(err)
	}
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	if err := os.Mkdir(wd+"/index/", 0700); err != nil {
		if !os.IsExist(err) { // Don't exit if index directory already exists.
			log.Fatal(err)
		}
	}
	for i := 1; i <= lastComic.Num; i++ {
		f, err := os.OpenFile(fmt.Sprintf("index/xkcd%d.json", i), os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			log.Fatal(err)
		}
		r, err := http.Get(fmt.Sprintf("https://xkcd.com/%d/info.0.json", i))
		if err != nil {
			f.Close()
			log.Fatal(err)
		}
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			r.Body.Close()
			f.Close()
			log.Fatal(err)
		}
		r.Body.Close()
		f.WriteString(string(b))
		f.Close()
	}
	fmt.Println("COMPLETE!")
}

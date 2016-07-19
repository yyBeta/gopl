package main

import (
	"fmt"
	"gopl/5函数/links"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{} // acquire a token
	list, err := links.Extract(url)
	<-tokens // release the token
	if err != nil {
		log.Print(err)
	}

	list = selectSameDomain(url, list)
	for _, link := range list {
		save(link)
	}

	return list
}

// setlect urls in the same domain
func selectSameDomain(link string, list []string) []string {
	srcURL, err := url.Parse(link)
	if err != nil {
		log.Print(err)
		return nil
	}
	filted := []string{}
	for _, l := range list {
		crawlURL, err := url.Parse(l)
		if err != nil {
			log.Print(err)
			return nil
		}
		if crawlURL.Host == srcURL.Host {
			filted = append(filted, l)
		}
	}
	return filted
}

// save the content of url
func save(link string) {
	resp, err := http.Get(link)
	if err != nil {
		log.Print(err)
		return
	}
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Print(err)
		return
	}
	theURL, err := url.Parse(link)
	if err != nil {
		log.Print(err)
		return
	}
	dir := filepath.Join("./", theURL.Host, filepath.Clean(theURL.Path))
	filename := "content"

	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		log.Print(err)
		return
	}

	err = ioutil.WriteFile(filepath.Join(dir, filename), b, os.ModePerm)
	if err != nil {
		log.Print(err)
		return
	}

}

func main() {
	worklist := make(chan []string)
	var n int // number of pending sends to worklist

	// Start with the command-line arguments.
	n++
	go func() { worklist <- os.Args[1:] }()

	// Crawl the web concurrently.
	seen := make(map[string]bool)

	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}

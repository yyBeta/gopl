package main

import (
	"fmt"
	"gopl/8Goroutines和Channels/links" //modified from ch5/links
	"log"
	"os"
)

var tokens = make(chan struct{}, 20)

func crawl(url string, cancel <-chan struct{}) []string {
	fmt.Println(url)
	tokens <- struct{}{} // acquire a token
	list, err := links.Extract(url, cancel)
	<-tokens // release the token
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	worklist := make(chan []string)
	cancel := make(chan struct{})
	var n int // number of pending sends to worklist

	// Start with the command-line arguments.
	n++
	go func() { worklist <- os.Args[1:] }()
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		close(cancel)
	}()

	// Crawl the web concurrently.
	seen := make(map[string]bool)

	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					worklist <- crawl(link, cancel)
				}(link)
			}
		}
	}
}

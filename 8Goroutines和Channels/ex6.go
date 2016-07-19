package main

import (
	"flag"
	"fmt"
	"gopl/5函数/links"
	"log"
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
	return list
}

//var depth = flag.Int("depth", 2, "max crawl depth. Default:2.")

type work struct {
	link  []string
	depth int
}

func main() {
	var depth int
	flag.IntVar(&depth, "depth", 2, "max crawl depth")
	flag.Parse()
	worklist := make(chan work)
	var n int // number of pending sends to worklist

	// Start with the command-line arguments.
	n++
	fmt.Println("max crawl depth :", depth)
	go func() { worklist <- work{flag.Args(), 0} }()

	// Crawl the web concurrently.
	seen := make(map[string]bool)

	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list.link {
			if !seen[link] && list.depth <= depth {
				seen[link] = true
				n++
				go func(link string, dep int) {
					worklist <- work{crawl(link), dep + 1}
				}(link, list.depth)
			}
		}
	}
}

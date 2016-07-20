package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

func fetch(urls []string) string {
	responses := make(chan string, len(urls))
	done := make(chan struct{})
	for _, u := range urls {
		go func(url string) {
			req, err := http.NewRequest("GET", url, nil)
			if err != nil {
				log.Printf("make request error %s", err)
				return
			}
			req.Cancel = done
			if _, err := http.DefaultClient.Do(req); err != nil {
				log.Printf("request failed %s", err)
				return
			}
			done <- struct{}{}
			responses <- url
		}(u)
	}
	return <-responses // return the quickest response
}

func main() {
	log.Printf("fastest URL is %s", fetch(os.Args[1:]))
	time.Sleep(3 * time.Second)
}

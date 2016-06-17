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

// breadthFirst calls f for each item in the worklist.
// Any items returned by f are added to the worklist.
// f is called at most once for each item.
func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func crawl(theURL string) []string {
	fmt.Println(theURL)
	list, err := links.Extract(theURL)
	if err != nil {
		log.Print(err)
	}
	list = selectSameDomain(theURL, list)

	for _, link := range list {
		save(link)
	}

	return list
}

// 筛选出相同域名的网址
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

// 保存该条地址对应的内容
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
	// Crawl the web breadth-first,
	// starting from the command-line arguments.
	breadthFirst(crawl, os.Args[1:])
}

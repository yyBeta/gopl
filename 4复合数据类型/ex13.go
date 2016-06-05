package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

const titleURL = "https://www.omdbapi.com/?t="

type movie struct {
	Response string
	Error    string
	Title    string
	Year     string
	Poster   string
}

func main() {
	query := url.QueryEscape(strings.Join(os.Args[1:], " "))
	if query == "" {
		fmt.Print("Error: No query provided.")
		os.Exit(1)
	}
	fmt.Printf("Querying Open Movie Database for title %#v...\n", query)
	resp, err := http.Get(titleURL + query)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	var myMovie movie
	if err := json.NewDecoder(resp.Body).Decode(&myMovie); err != nil {
		log.Fatal(err)
	}
	if myMovie.Response == "False" { // Movie not found!
		fmt.Print("Error: " + myMovie.Error)
		os.Exit(1)
	}
	fmt.Printf("Found movie %#v (%s)\n", myMovie.Title, myMovie.Year)
	if myMovie.Poster == "N/A" {
		fmt.Print("Error: No poster image available.")
		os.Exit(1)
	}
	fmt.Print("Downloading movie poster... ")
	resp, err = http.Get(myMovie.Poster)
	if err != nil {
		log.Fatal(err)
	}
	filename := filepath.Base(myMovie.Poster)
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := io.Copy(file, resp.Body); err != nil {
		file.Close()
		os.Remove(filename)
		log.Fatal(err)
	}
	file.Close()
	fmt.Println("Success!")
}

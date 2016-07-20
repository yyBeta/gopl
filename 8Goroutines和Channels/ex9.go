package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var vFlag = flag.Bool("v", false, "show verbose progress messages")

type sizeResponse struct {
	root int
	size int64
}

func main() {
	// ...determine roots...

	flag.Parse()

	// Determine the initial directories.
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	// Traverse each root of the file tree in parallel.
	sizeResponses := make(chan sizeResponse)
	var n sync.WaitGroup
	for i, root := range roots {
		n.Add(1)
		go walkDir(root, i, &n, sizeResponses)
	}
	go func() {
		n.Wait()
		close(sizeResponses)
	}()

	// Print the results periodically.
	var tick <-chan time.Time
	if *vFlag {
		tick = time.Tick(1000 * time.Millisecond)
	}
	nfiles := make([]int64, len(roots))
	nbytes := make([]int64, len(roots))

loop:
	for {
		select {
		case sr, ok := <-sizeResponses:
			if !ok {
				break loop // fileSizes was closed
			}
			nfiles[sr.root]++
			nbytes[sr.root] += sr.size
		case <-tick:
			printDiskUsage(roots, nfiles, nbytes)
		}
	}

	printDiskUsage(roots, nfiles, nbytes) // final totals
	// ...select loop...
}

func printDiskUsage(roots []string, nfiles, nbytes []int64) {
	for i, r := range roots {
		fmt.Printf("%d files  %.2f GB under %s\n", nfiles[i], float64(nbytes[i])/1e9, r)
	}
}

// walkDir recursively walks the file tree rooted at dir
// and sends the size of each found file on fileSizes.

func walkDir(dir string, i int, n *sync.WaitGroup, sr chan<- sizeResponse) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, i, n, sr)
		} else {
			sr <- sizeResponse{i, entry.Size()}
		}
	}
}

// sema is a counting semaphore for limiting concurrency in dirents.
var sema = make(chan struct{}, 20)

// dirents returns the entries of directory dir.
func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}        // acquire token
	defer func() { <-sema }() // release token
	// ...

	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	return entries
}

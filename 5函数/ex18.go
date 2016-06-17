package main // Fetch downloads the URL and returns the
import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

// name and length of the local file.
func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	defer func() {
		if closeErr := f.Close(); err == nil {
			err = closeErr
		}
	}()

	n, err = io.Copy(f, resp.Body)
	return local, n, err
}

func main() {
	fmt.Println(fetch("http://gopl.io"))
}

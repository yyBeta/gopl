package main // WaitForServer attempts to contact the server of a URL.
import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

// 1.最常见的是传播错误，传给调用者

func WaitForServer(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	// 2.重新尝试失败的操作。在重试时，我们需要限制重试的时间间隔或重试的次数，防止无限制的重试
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil // success
		}
		log.Printf("server not responding (%s);retrying…", err)
		time.Sleep(time.Second << uint(tries)) // exponential back-off
	}
	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}

func main() {
	// 3.输出错误信息并结束程序
	if err := WaitForServer(url); err != nil {
		log.Fatalf("Site is down: %v\n", err)
		//the same as: fmt.Fprintf(os.Stderr, "Site is down: %v\n", err)
		os.Exit(1)
	}

	// 4.仅输出错误
	if err := Ping(); err != nil {
		log.Printf("ping failed: %v; networking disabled", err)
		//the same as: fmt.Fprintf(os.Stderr, "ping failed: %v; networking disabled\n", err)
	}

	// 5.直接忽略掉错误
	dir, err := ioutil.TempDir("", "scratch")
	if err != nil {
		return fmt.Errorf("failed to create temp dir: %v", err)
	}
	// ...use temp dir…
	os.RemoveAll(dir) // ignore errors; $TMPDIR is cleaned periodically

}

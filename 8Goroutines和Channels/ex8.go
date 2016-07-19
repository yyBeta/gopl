package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	var wg sync.WaitGroup
	reset := make(chan struct{})
	input := bufio.NewScanner(c)
	go func() {
		for input.Scan() {
			reset <- struct{}{} // send time reseting signal if anything scaned
			wg.Add(1)
			go func(s string) {
				defer wg.Done()
				echo(c, s, 1*time.Second)
			}(input.Text())
		}
	}()
	// NOTE: ignoring potential errors from input.Err()
	for {
		select {
		case <-time.After(10 * time.Second):
			wg.Wait() // do not close connection before echos done
			if tc, ok := c.(*net.TCPConn); ok {
				tc.CloseWrite()
			} else {
				c.Close()
			}
			fmt.Println("long time no input, connection closed.")
			return
		case <-reset:
			// do nothing
		}
	}
}

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}
}

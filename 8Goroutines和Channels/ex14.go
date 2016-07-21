package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

type client struct {
	msg chan<- string // an outgoing message channel
	who string
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // all incoming client messages
)

func broadcaster() {
	clients := make(map[client]bool) // all connected clients
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			for cli := range clients {
				cli.msg <- msg
			}
		case cli := <-entering:
			clients[cli] = true
			cli.msg <- func() string {
				whos := make([]string, len(clients))
				for cli := range clients {
					whos = append(whos, cli.who)
				}
				return fmt.Sprintf("current clients: %v", whos)
			}()

		case cli := <-leaving:
			delete(clients, cli)
			close(cli.msg)
		}
	}
}

func handleConn(conn net.Conn) {
	fmt.Fprintf(conn, "Enter your name: ")
	input := bufio.NewScanner(conn)
	input.Scan()
	who := input.Text() // ignore inputing errors

	ch := make(chan string) // outgoing client messages
	go clientWriter(conn, ch)

	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- client{ch, who}

	reset := make(chan struct{})
	end := make(chan struct{})
	go func() {
		for input.Scan() {
			reset <- struct{}{}
			messages <- who + ": " + input.Text()
		}
		end <- struct{}{}
	}()
	// NOTE: ignoring potential errors from input.Err()
loop:
	for {
		select {
		case <-time.After(5 * time.Minute):
			break loop
		case <-end:
			break loop
		case <-reset:
			// do nothing
		}
	}
	leaving <- client{ch, who}
	messages <- who + " has left"
	conn.Close()

}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}

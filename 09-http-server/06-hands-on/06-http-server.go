package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	var name string
	fmt.Print("what's your name? ")
	fmt.Scanln(&name)
	fmt.Printf("Hello %v\n", name)

	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer li.Close()

	fmt.Println("Listening on localhost:8080")
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		go httpHandler(conn)
	}
}

func httpHandler(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		li := scanner.Text()
		if li == "" {
			break
		}
	}
	resp := `HTTP/1.1 302 Found
Location: https://www.google.com
`
	fmt.Fprintf(conn, resp)
	fmt.Println("Closing the connection")
}

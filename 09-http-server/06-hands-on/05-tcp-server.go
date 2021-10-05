package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
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
		go tcpHandler(conn)
	}
}

func tcpHandler(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		li := scanner.Text()
		fmt.Println(li)
		fmt.Fprintf(conn, "I got: %s\n", li)
	}
	fmt.Println("Closing the connection")
}

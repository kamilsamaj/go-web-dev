package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln(err)
		}

		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	err := conn.SetDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		fmt.Println("CONN TIMEOUT")
	}
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		fmt.Fprintf(conn, "I heard you say: %s\n", ln)
	}
	//buf := make([]byte, 4096)
	//fmt.Print(conn, "What's your name? ")
	//_, err = conn.Read(buf)
	//if err != nil {
	//	log.Println(err)
	//}
	//fmt.Println("I heard you saying:", string(buf))
	//fmt.Fprintf(conn, string(buf))
	fmt.Println("code got here")
}

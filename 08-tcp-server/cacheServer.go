package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

// share cache for all clients
var cache map[string]string

func main() {
	li, err := net.Listen("tcp", ":8080")
	cache = make(map[string]string)
	if err != nil {
		panic(err)
	}
	defer li.Close()
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		go handler(conn)
	}

}

func handler(conn net.Conn) {
	defer conn.Close()
	// initial usage help
	io.WriteString(conn, `IN-MEMORY DATABASE

USE:
	SET key value
	GET key
	DEL key

EXAMPLE:
	SET fav chocolate
	GET fav
	DEL fav

> `)

	// create

	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		ln := scanner.Text()
		fs := strings.Fields(ln)

		switch strings.ToUpper(fs[0]) {
		case "GET":
			k := fs[1]
			v := cache[k]
			fmt.Fprintf(conn, "%s\n", v)
		case "SET":
			if len(fs) != 3 {
				fmt.Fprintln(conn, "Error: Usage: SET key value\n")
				continue
			}
			k := fs[1]
			v := fs[2]
			cache[k] = v
		case "DEL":
			delete(cache, fs[1])
		default:
			fmt.Println(conn, "INVALID COMMAND:", fs[0], "\n")
			continue
		}
		// display prompt
		fmt.Fprint(conn, "> ")
	}
}

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
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		go handler(conn)
	}
}

func handler(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	defer conn.Close()
	var reqLines []string
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if ln == "" {
			break // browsers dont close connections
		}
		reqLines = append(reqLines, ln)
	}
	//
	//for _, ln := range reqLines {
	//	io.WriteString(conn, ln)
	//}

	body := `<!doctype html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport"
          content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Pure TCP response</title>
</head>
<body>
<h1>Hello from TCP</h1>
</body>
</html>`

	fmt.Fprintln(conn, "HTTP/1.1 200 OK")
	fmt.Fprintf(conn, "Content-Length: %d\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\n\n")
	fmt.Fprint(conn, body)
}

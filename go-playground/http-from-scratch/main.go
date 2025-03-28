package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func main() {
	listen, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatal("Error occurred starting the server: ", err)
	}

	fmt.Println("Server listening on http://localhost:8080")

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Failed to accept connection: ", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 1024)

	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err != io.EOF {
				log.Println("Error reading from a file: ", err)
			}
			break
		}

		rawRequest := string(buf[:n])
		lines := strings.Split(rawRequest, "\r\n")
		requestLine := lines[0]
		parts := strings.Split(requestLine, " ")

		if len(parts) == 3 {
			method := parts[0]
			path := parts[1]
			version := parts[2]

			fmt.Println("Method: " + method)
			fmt.Println("Path: " + path)
			fmt.Println("Version: " + version)
		}

		if n > 0 {
			fmt.Println("Received HTTP request:")
			fmt.Println(rawRequest)

			body := "Hello from antdevrealm"
			response := fmt.Sprintf("HTTP/1.1 200 OK\r\n"+
				"Content-Type: text/plain\r\n"+
				"Content-Length: %d\r\n"+
				"\r\n%s",
				len(body), body,
			)

			_, err := conn.Write([]byte(response))

			if err != nil {
				log.Println("Error writing message: ", err)
			}
			break
		}
	}
}

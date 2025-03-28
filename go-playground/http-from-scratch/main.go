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
		var method, path, version string

		var body string
		var status string
		var response string

		if len(lines) > 0 {
			requestLine := lines[0]
			parts := strings.Split(requestLine, " ")

			if len(parts) == 3 {
				method = parts[0]
				path = parts[1]
				version = parts[2]
			}
		}

		if method == "" || path == "" || version == "" {
			status = "400 Bad Request"
			body = "400 Bad Request"
			response = fmt.Sprintf(
				"HTTP/1.1 %s\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s",
				status, len(body), body,
			)
			conn.Write([]byte(response))
			log.Println("Bad Request!")
			return
		}

		if method == "GET" {
			switch path {
			case "/":
				status = "200 OK"
				body = "Welcome to the homepage"
				break

			case "/about":
				status = "200 OK"
				body = "This is a raw Go TCP server built with LOVE"

			default:
				status = "404 Not Found"
				body = "404 Not Found"
			}
		}

		if status != "" {
			response = fmt.Sprintf(
				"HTTP/1.1 %s\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s",
				status, len(body), body,
			)

			_, err := conn.Write([]byte(response))
			if err != nil {
				log.Println("Error writing response: ", err)
			}
		}

	}
}

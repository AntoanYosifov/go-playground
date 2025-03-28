package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}
	fmt.Println("Server listening on http://localhost:8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Failed to accept connection:", err)
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
				log.Println("Error reading from connection:", err)
			}
			break
		}
		if n > 0 {
			fmt.Println("📥 Received HTTP request:")
			fmt.Println(string(buf[:n]))

			response := "HTTP/1.1 200 OK\r\n" +
				"Content-Type: text/plain\r\n" +
				"Content-Length: 16\r\n" +
				"\r\n" +
				"Hello from Go!"

			_, err := conn.Write([]byte(response))
			if err != nil {
				log.Println("Error writing response:", err)
			}
			break
		}
	}
}

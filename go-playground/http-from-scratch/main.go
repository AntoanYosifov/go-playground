package main

import (
	"fmt"
	"http-from-scratch/models"
	"http-from-scratch/parser"
	"http-from-scratch/response"
	"http-from-scratch/router"
	"io"
	"log"
	"net"
)

var jobs []models.Job = []models.Job{}

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
		req := parser.Parse(rawRequest)

		var body string
		var status string
		var job *models.Job
		var responseValue string

		if req.Method == "" || req.Path == "" || req.Version == "" {
			status = "400 Bad Request"
			body = "400 Bad Request"
			responseValue = response.BuildResponse(status, body)
			conn.Write([]byte(responseValue))
			log.Println("Bad Request")
			return
		}

		if req.Method == "GET" {
			status, body = router.HandleGet(req.Path, jobs)
		}

		if req.Method == "POST" {
			status, body, job = router.HandlePost(req.Path, req.Body, req.Headers)

			if job != nil {
				jobs = append(jobs, *job)
			}
		}

		if status != "" {
			responseValue = response.BuildResponse(status, body)
			_, err := conn.Write([]byte(responseValue))
			if err != nil {
				log.Println("Error writing response: ", err)
			}
		}
	}
}

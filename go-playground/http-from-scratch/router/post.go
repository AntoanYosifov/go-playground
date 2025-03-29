package router

import (
	"encoding/json"
	"fmt"
	"http-from-scratch/models"
	"strings"
)

func HandlePost(path string, body string, headers map[string]string) (status string, response string, job *models.Job) {
	switch path {
	case "/jobs":
		contentType := headers["Content-Type"]

		if strings.Contains(contentType, "application/json") {
			var job models.Job
			err := json.Unmarshal([]byte(body), &job)
			if err != nil {
				return "400 Bad Request", "Invalid JSON", nil
			}

			fmt.Println("Job received: ", job)
			return "201 Created", "Job added", &job

		} else if strings.Contains(contentType, "application/x-www-form-urlencoded") {
			return "200 OK", "Received Form:\n" + body, nil
		}

		return "415 Unsupported Media Type", "Unsupported Content-Type: " + contentType, nil

	default:
		return "404 Not Found", "404 Not Found", nil
	}
}

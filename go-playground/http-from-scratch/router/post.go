package router

import (
	"encoding/json"
	"fmt"
	"http-from-scratch/models"
	"strings"
)

func HandlePost(path string, body string, headers map[string]string, jobs *[]models.Job) (status string, response string) {
	switch path {
	case "/jobs":
		contentType := headers["Content-Type"]

		if strings.Contains(contentType, "application/json") {
			var newJobs []models.Job
			err := json.Unmarshal([]byte(body), &newJobs)
			if err != nil {
				return "400 Bad Request", "Invalid JSON array"
			}

			*jobs = append(*jobs, newJobs...)
			return "201 Created", fmt.Sprintf("Added %d job(s)", len(newJobs))

		} else if strings.Contains(contentType, "application/x-www-form-urlencoded") {
			return "200 OK", "Received Form:\n" + body
		}

		return "415 Unsupported Media Type", "Unsupported Content-Type: " + contentType

	default:
		return "404 Not Found", "404 Not Found"
	}
}

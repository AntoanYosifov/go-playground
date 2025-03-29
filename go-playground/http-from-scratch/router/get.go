package router

import (
	"encoding/json"
	"http-from-scratch/models"
)

func HandleGet(path string, jobs []models.Job) (string, string) {
	switch path {
	case "/":
		return "200 OK", "Welcome to the homepage"

	case "/about":
		return "200 OK", "This is a raw Go server build with LOVE!"

	case "/jobs":
		data, err := json.MarshalIndent(jobs, "", "  ")
		if err != nil {
			return "500 Internal Server Error", "Failed to encode jobs"
		}

		return "200 OK", string(data)

	default:
		return "404 Not Found", "404 Not Found"
	}

}

package router

import "strings"

func HandlePost(path string, body string, headers map[string]string) (status, response string) {
	switch path {
	case "/submit":
		contentType := headers["Content-Type"]

		if strings.Contains(contentType, "application/json") {
			return "200 OK", "Received JSON:\n" + body
		} else if strings.Contains(contentType, "application/x-www-form-urlencoded") {
			return "200 OK", "Received Form:\n" + body
		}

		return "415 Unsupported Media Type", "Unsupported Content-Type: " + contentType

	default:
		return "404 Not Found", "404 Not Found"
	}
}

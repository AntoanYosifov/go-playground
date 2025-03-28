package response

import "fmt"

func BuildResponse(status, body string) string {
	return fmt.Sprintf(
		"HTTP/1.1 %s\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s",
		status, len(body), body,
	)
}

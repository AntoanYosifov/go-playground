package parser

import "strings"

type Request struct {
	Method  string
	Path    string
	Version string
	Headers map[string]string
	Body    string
}

func Parse(raw string) Request {
	lines := strings.Split(raw, "\r\n")
	req := Request{
		Headers: make(map[string]string),
	}

	if len(lines) > 0 {
		method, path, version := ParseRequestLine(lines[0])
		req.Method = method
		req.Path = path
		req.Version = version
	}

	for _, line := range lines[1:] {
		if line == "" {
			break
		}
		if parts := strings.SplitN(line, ": ", 2); len(parts) == 2 {
			req.Headers[parts[0]] = parts[1]
		}
	}

	bodyStart := strings.Index(raw, "\r\n\r\n")
	if bodyStart != -1 && bodyStart+4 < len(raw) {
		req.Body = raw[bodyStart+4:]
	}

	return req
}

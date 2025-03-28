package parser

import "strings"

func ParseRequestLine(line string) (method, path, version string) {
	parts := strings.Split(line, " ")
	if len(parts) == 3 {
		return parts[0], parts[1], parts[2]
	}
	return "", "", ""
 }

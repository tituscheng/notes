package util

import (
	"strings"
)

func GetTitle(content string) string {
	if len(content) == 0 {
		return "New Note"
	}
	lines := strings.Split(content, "\n")
	return lines[0]
}

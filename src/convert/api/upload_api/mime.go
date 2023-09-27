package uploadapi

import (
	"strings"
)

func ReturnMime(s string) string {
	parts := strings.Split(s, "/")
	var ext string
	if len(parts) > 1 {
		ext = parts[1]
		if parts[1] == "x-matroska" {
			ext = "mkv"
		}
	}
	
	return ext
}


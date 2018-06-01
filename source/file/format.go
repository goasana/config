package file

import (
	"strings"
)

func format(p string) string {
	parts := strings.Split(p, ".")
	if len(parts) > 1 {
		return parts[len(parts)-1]
	}
	return "json"
}

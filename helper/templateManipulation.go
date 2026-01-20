package helper

import (
	"strings"
)

func Truncate(s string, limit int) string {
	var sb strings.Builder
	for _, c := range s {
		if len(sb.String()) < limit {
			sb.WriteRune(c)
		}
	}
	sb.WriteString("...")
	return sb.String()
}

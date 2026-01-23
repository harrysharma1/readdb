package helper

import (
	"fmt"
	"readdb/models"
	"strings"
)

func Truncate(s string, limit int) string {
	var count int
	var sb strings.Builder
	for _, c := range s {
		if count >= limit {
			break
		}
		count++
		sb.WriteRune(c)
	}
	sb.WriteString("...")
	return sb.String()
}

func Linkify(bookBuyLink models.BookBuyLink) string {
	switch bookBuyLink.Retailer {
	case "Waterstones":
		retVal := fmt.Sprintf(`<a href="%s" target="_blank" class="btn btn-default">Buy at Waterstones</a>`, bookBuyLink.URL)
		return retVal
	case "Amazon":
		return fmt.Sprintf(`<a href="%s" target="_blank" class="btn btn-default">Buy at Amazon</a>`, bookBuyLink.URL)
	default:
		return ""
	}
}

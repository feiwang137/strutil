package strutil

import (
	"fmt"
	"strings"
)

func EscapeHTMLTag(html string) string {
	ESCAPE_HTML := map[string]string{
		"&":  "&amp;",
		"<":  "&lt;",
		">":  "&gt;",
		"'":  "&#39;",
		"\"": "&quot;",
	}

	result := html

	for k, v := range ESCAPE_HTML {
		result = strings.ReplaceAll(result, k, v)
	}
	
	fmt.Println("...... Debug v0.4.1......")
	
	return result
}

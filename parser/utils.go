package main

import (
	"fmt"
	"strings"
)

func capitalize(key string) string {
	return fmt.Sprintf("%s%s", strings.ToUpper(string(key[0])), key[1:])
}

func formatRange(r []Range) string {
	if len(r) == 1 {
		if r[0].Min == r[0].Max {
			if r[0].Min == 1 { // 1
				return ""
			}
			return fmt.Sprintf("[%d]", r[0].Min) // n
		}
		if r[0].Min == 0 {
			if r[0].Max == 1 { // 0 or 1, 0 to 1
				return "*"
			}
		}
		// 0 or ., 0 to ., 0 to n
		return "*[]"
	}
	// Nothing is interpreted as 1
	return ""
}

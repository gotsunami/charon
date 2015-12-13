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
			if r[0].Min == 1 {
				return ""
			}
			return fmt.Sprintf("[%d]", r[0].Min)
		}
		if r[0].Min == 0 {
			if r[0].Max == 1 {
				return "*"
			}
		}
		return "[]"
	}
	return ""
}

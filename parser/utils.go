package main

import (
	"fmt"
	"strings"
)

func capitalize(key string) string {
	return fmt.Sprintf("%s%s", strings.ToUpper(string(key[0])), key[1:])
}

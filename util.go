package main

import (
	"strings"
)

func endWith(s, keyword string) bool {
	return len(s)-len(keyword) == strings.LastIndex(s, keyword)
}

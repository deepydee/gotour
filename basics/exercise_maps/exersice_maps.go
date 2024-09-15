package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	result := make(map[string]int)

	src := strings.Fields(s)
	for _, v := range src {
		result[v]++
	}

	return result
}

func main() {
	wc.Test(WordCount)
}

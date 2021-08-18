package main

import (
	"fmt"
	"strings"
)

// import (
// 	"golang.org/x/tour/wc"
// )

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	counts := make(map[string]int)
	for i := 0; i < len(words); i++ {
		counts[strings.Trim(words[i], ".,")] += 1
	}
	return counts
}

func main() {
	// wc.Test(WordCount)
	fmt.Println(WordCount("Here are some words, words, words, and more words for more fun are we having here."))
}

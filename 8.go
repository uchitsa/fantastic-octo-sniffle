package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)

	for i := 0; i < t; i++ {
		var n, b, r, f int
		fmt.Fscan(in, &n, &b, &r, &f)

		words := make([]string, n)
		for j := 0; j < n; j++ {
			fmt.Fscan(in, &words[j])
		}

		blackWord := words[:f-1]
		blueWords := words[:b]
		redWords := words[b : b+r]

		bestWord := ""
		bestDiff := -1

		for j := 0; j < n; j++ {
			if j == f-1 { // Skip the black word
				continue
			}

			word := words[j]
			if isSubstringOfAny(word, blackWord) {
				continue
			}

			blueCount := countSubstringsInList(word, blueWords)
			redCount := countSubstringsInList(word, redWords)
			diff := blueCount - redCount

			if diff > bestDiff {
				bestDiff = diff
				bestWord = word
			}
		}

		fmt.Fprintln(out, bestWord, bestDiff)
	}
}

func isSubstringOfAny(word string, words []string) bool {
	for _, w := range words {
		if isSubstring(word, w) {
			return true
		}
	}
	return false
}

func isSubstring(sub string, s string) bool {
	n := len(sub)
	for i := 0; i <= len(s)-n; i++ {
		if s[i:i+n] == sub {
			return true
		}
	}
	return false
}

func countSubstringsInList(word string, words []string) int {
	count := 0
	for _, w := range words {
		if isSubstring(word, w) {
			count++
		}
	}
	return count
}

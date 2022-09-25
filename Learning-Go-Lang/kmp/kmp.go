package main

import "fmt"

func createLpsTable(pattern string) []int {
	longestPrefixSuffixTable := make([]int, len(pattern)) //slice; size of the pattern string

	for i, j := 0, 1; j < len(pattern); {
		// fmt.Println("hehe", i)
		switch {
		case pattern[j] == pattern[i]:
			longestPrefixSuffixTable[j] = i + 1
			i++
			j++
		case i == 0:
			j++ // search for match -> go next char
		default:
			i = longestPrefixSuffixTable[i-1] //return to the last matched position
		}
	}

	return longestPrefixSuffixTable
}

func kmp(s string, pattern string) int {

	lps := createLpsTable(pattern)
	matchedAt := 0

	for i, j := 0, 0; i <= len(s); { //i to trace string s, j to trace pattern
		if j == len(pattern) {
			matchedAt = i - len(pattern)
			break
		}

		if i == len(s) { //end of string s
			break
		}

		if s[i] == pattern[j] {
			i++ //to check if next char is same => string s
			j++ //pattern
		} else {
			if j == 0 { //no char of pattern is matched yet
				i++ //as no char matched, check with next char of s
			} else { // j != 0
				j = lps[j-1]
			}
		}
	}

	return matchedAt
}

func main() {
	fmt.Print("Pattern matches at index: ")
	// fmt.Print(strStr("ABABDABACDABABCABAB", "ABABCABAB"))
	fmt.Print(kmp("BBBBBBBBBBBBBBBBBBBAB", "AB"))
}

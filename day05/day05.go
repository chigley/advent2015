package day05

import (
	"strings"
)

var vowels = map[rune]struct{}{
	'a': {},
	'e': {},
	'i': {},
	'o': {},
	'u': {},
}

var bad = []string{"ab", "cd", "pq", "xy"}

func Part1(in []string) int {
	return totalNice(in, isNicePart1)
}

func totalNice(in []string, isNice func(s string) bool) int {
	var numNice int
	for _, s := range in {
		if isNice(s) {
			numNice++
		}
	}
	return numNice
}

func isNicePart1(s string) bool {
	for _, b := range bad {
		if strings.Contains(s, b) {
			return false
		}
	}

	var (
		numVowels  int
		haveDouble bool
	)

	for i, c := range s {
		if _, ok := vowels[c]; ok {
			numVowels++
		}

		if i > 0 && rune(s[i-1]) == c {
			haveDouble = true
		}
	}

	return haveDouble && numVowels >= 3
}

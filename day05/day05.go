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

func Part2(in []string) int {
	return totalNice(in, isNicePart2)
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

func isNicePart2(s string) bool {
	startIndices := make(map[string][]int)

	var haveSeparatedSinglePair bool
	for i, c := range s {
		if i != len(s)-1 {
			pair := s[i : i+2]
			startIndices[pair] = append(startIndices[pair], i)
		}

		if i > 1 && rune(s[i-2]) == c {
			haveSeparatedSinglePair = true
		}
	}

	if !haveSeparatedSinglePair {
		return false
	}

	for _, indices := range startIndices {
		// If we have 3 pairs then at least two don't overlap.
		if len(indices) >= 3 {
			return true
		}

		// If we have 2 pairs, make sure they don't overlap.
		if len(indices) == 2 && indices[1]-indices[0] >= 2 {
			return true
		}
	}

	return false
}

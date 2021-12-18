package day01

import (
	"fmt"

	"github.com/chigley/advent2015"
)

func Part1(in string) (int, error) {
	var floor int
	for _, c := range in {
		switch c {
		case '(':
			floor++
		case ')':
			floor--
		default:
			return 0, fmt.Errorf("day01: unexpected %c", c)
		}
	}
	return floor, nil
}

func Part2(in string) (int, error) {
	var floor int
	for i, c := range in {
		switch c {
		case '(':
			floor++
		case ')':
			floor--
		default:
			return 0, fmt.Errorf("day01: unexpected %c", c)
		}

		if floor == -1 {
			return i + 1, nil
		}
	}
	return 0, advent2015.ErrNoSolution
}

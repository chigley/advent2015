package day01

import (
	"fmt"
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

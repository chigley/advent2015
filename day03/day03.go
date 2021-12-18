package day03

import (
	"fmt"

	"github.com/chigley/advent2015"
)

type Grid map[advent2015.XY]int

func Part1(in string) (int, error) {
	var pos advent2015.XY

	g := Grid{
		pos: 1,
	}

	for _, c := range in {
		switch c {
		case '^':
			pos = pos.Add(advent2015.North)
		case 'v':
			pos = pos.Add(advent2015.South)
		case '>':
			pos = pos.Add(advent2015.East)
		case '<':
			pos = pos.Add(advent2015.West)
		default:
			return 0, fmt.Errorf("day03: invalid direction %c", c)
		}

		g[pos]++
	}

	return len(g), nil
}

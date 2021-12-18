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

func Part2(in string) (int, error) {
	g := Grid{
		advent2015.XY{X: 0, Y: 0}: 2,
	}

	var santa, roboSanta advent2015.XY
	for i, c := range in {
		var pos *advent2015.XY
		if i%2 == 0 {
			pos = &santa
		} else {
			pos = &roboSanta
		}

		switch c {
		case '^':
			*pos = pos.Add(advent2015.North)
		case 'v':
			*pos = pos.Add(advent2015.South)
		case '>':
			*pos = pos.Add(advent2015.East)
		case '<':
			*pos = pos.Add(advent2015.West)
		default:
			return 0, fmt.Errorf("day03: invalid direction %c", c)
		}

		g[*pos]++
	}

	return len(g), nil
}

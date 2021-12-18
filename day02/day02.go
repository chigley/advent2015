package day02

import (
	"fmt"

	"github.com/chigley/advent2015"
)

type Dimensions struct {
	L, W, H int
}

func Part1(ds []Dimensions) int {
	var total int
	for _, d := range ds {
		total += d.wrappingPaper()
	}
	return total
}

func Part2(ds []Dimensions) int {
	var total int
	for _, d := range ds {
		total += d.ribbon()
	}
	return total
}

func (d Dimensions) wrappingPaper() int {
	sides := []int{
		d.L * d.W,
		d.W * d.H,
		d.H * d.L,
	}

	paper := advent2015.MinInts(sides)
	for _, n := range sides {
		paper += 2 * n
	}
	return paper
}

func (d Dimensions) ribbon() int {
	sides := []int{
		d.L + d.W,
		d.W + d.H,
		d.H + d.L,
	}

	return 2*advent2015.MinInts(sides) + d.L*d.W*d.H
}

func ParseDimensions(in []string) ([]Dimensions, error) {
	ds := make([]Dimensions, len(in))
	for i, l := range in {
		var d Dimensions
		if _, err := fmt.Sscanf(l, "%dx%dx%d", &d.L, &d.W, &d.H); err != nil {
			return nil, fmt.Errorf("day02: parsing %q: %w", l, err)
		}
		ds[i] = d
	}
	return ds, nil
}

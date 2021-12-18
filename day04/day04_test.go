package day04_test

import (
	"testing"

	"github.com/chigley/advent2015/day04"
	"github.com/stretchr/testify/assert"
)

const input = "yzbqklnj"

var tests = []struct {
	in    string
	part1 int
	part2 int
}{
	{"abcdef", 609043, 6742839},
	{"pqrstuv", 1048970, 5714438},
	{input, 282749, 9962624},
}

func TestDay04(t *testing.T) {
	for _, tt := range tests {
		tt := tt

		t.Run(tt.in, func(t *testing.T) {
			t.Parallel()

			part1, err := day04.Part1(tt.in)
			if err != nil {
				t.Error(err)
			}

			part2, err := day04.Part2(tt.in)
			if err != nil {
				t.Error(err)
			}

			assert.Equal(t, tt.part1, part1)
			assert.Equal(t, tt.part2, part2)
		})
	}
}

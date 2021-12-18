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
}{
	{"abcdef", 609043},
	{"pqrstuv", 1048970},
	{input, 282749},
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

			assert.Equal(t, tt.part1, part1)
		})
	}
}

package day05_test

import (
	"testing"

	"github.com/chigley/advent2015"
	"github.com/chigley/advent2015/day05"
	"github.com/stretchr/testify/assert"
)

func TestDay05(t *testing.T) {
	t.Parallel()

	in, err := advent2015.ReadStrings("testdata/input")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 238, day05.Part1(in))
}

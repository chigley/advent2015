package day02_test

import (
	"testing"

	"github.com/chigley/advent2015/day02"
	"github.com/stretchr/testify/assert"

	"github.com/chigley/advent2015"
)

func TestDay02(t *testing.T) {
	t.Parallel()

	in, err := advent2015.ReadStrings("testdata/input")
	if err != nil {
		t.Fatal(err)
	}

	ds, err := day02.ParseDimensions(in)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 1598415, day02.Part1(ds))
}

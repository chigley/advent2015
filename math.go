package advent2015

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func MinInts(ns []int) int {
	if len(ns) == 0 {
		panic("min of empty slice")
	}

	min := ns[0]
	for _, n := range ns[1:] {
		min = Min(min, n)
	}
	return min
}

package advent2015

var (
	North = XY{0, 1}
	East  = XY{1, 0}
	South = XY{0, -1}
	West  = XY{-1, 0}
)

type XY struct {
	X, Y int
}

func (p XY) Add(q XY) XY {
	return XY{
		X: p.X + q.X,
		Y: p.Y + q.Y,
	}
}

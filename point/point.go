package point

import "github.com/bomgar/adventhelper/math"

type Point struct {
	Y, X int
}

func ManhattanDistance(p1, p2 Point) int {
	return math.Abs(p1.X-p2.X) + math.Abs(p1.Y-p2.Y)
}

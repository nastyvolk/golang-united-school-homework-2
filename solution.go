package square

import "math"

type intCustomType int

const (
	SidesTriangle intCustomType = 3
	SidesSquare   intCustomType = 4
	SidesCircle   intCustomType = 0
)

func CalcSquare(sideLen float64, sidesNum intCustomType) float64 {
	switch sidesNum {
	case SidesTriangle:
		return math.Sqrt(3) / 4 * math.Pow(sideLen, 2)
	case SidesSquare:
		return math.Pow(sideLen, 2)
	case SidesCircle:
		return math.Pi * math.Pow(sideLen, 2)
	default:
		return 0
	}
}

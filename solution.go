package square

import (
	"math"
)

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)
type sidesNum int

const (
	SidesSquare   sidesNum = 4
	SidesTriangle sidesNum = 3
	SidesCircle   sidesNum = 0
)

// CalcSquare computes area
func CalcSquare(sideLen float64, sidesNum sidesNum) float64 {
	sqrSideLen := sideLen * sideLen

	switch sidesNum {
	case SidesSquare:
		return sqrSideLen
	case SidesTriangle:
		return sqrSideLen * math.Sqrt(3) / 4
	case SidesCircle:
		return sqrSideLen * math.Pi
	default:
		return 0
	}

}

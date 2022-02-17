package square

import (
	"math"
)

type sidesType int

const SidesCircle sidesType = 0;
const SidesTriangle sidesType = 3;
const SidesSquare sidesType = 4;
// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

func CalcSquare(sideLen float64, sidesNum sidesType) float64 {
	var square float64

	switch sidesNum {
		case 0: square = calcCircle(sideLen)
		case 3: square = calcTringle(sideLen)
		case 4: square = calcSquare(sideLen)
	}

	return square
}

func calcTringle(sideLen float64) (square float64) {
	square = (math.Pow(sideLen, 2) * math.Sqrt(3)) / 4

	return
}

func calcSquare(sideLen float64) (square float64) {
	square = math.Pow(sideLen, 2)

	return
}

func calcCircle(sideLen float64) (square float64) {
	square = math.Pi * math.Pow(sideLen, 2)

	return
}

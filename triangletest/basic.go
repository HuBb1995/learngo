package triangletest

import (
	"fmt"
	"math"
)

func triangle() {
	fmt.Println(calcTriangle(3, 4))
}

func calcTriangle(a, b int) int {
	return int(math.Sqrt(float64(a*a + b*b)))
}

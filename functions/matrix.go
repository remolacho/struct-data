package functions

import (
	"fmt"
	"math"
)

func MatrixPow() {
	var m [3][9]int

	for i := 0; i < 3; i++ {
		for j := 0; j < 9; j++ {
			m[i][j] = int(math.Pow(float64(j+1), float64(i+1)))
			fmt.Printf("%d \t", m[i][j])
		}
		fmt.Printf("\n")
	}
}

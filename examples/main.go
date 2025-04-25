package main

import (
	"fmt"

	"github.com/phlucasfr/manhattan-distance/manhattan"
)

func main() {
	matrix := [][]int{
		{0, 0, 0, 0},
		{0, 1, 0, 0},
		{0, 0, 0, 1},
		{0, 0, 0, 0},
	}
	fmt.Println(manhattan.ManhattanDistance(matrix)) // Saída: 3
}

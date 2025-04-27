package main

import (
	"testing"

	"github.com/phlucasfr/manhattan-distance/manhattan"
)

func TestManhattanDistance(t *testing.T) {
	matrix := [][]int{
		{0, 0, 0, 0},
		{0, 1, 0, 0},
		{0, 0, 0, 1},
		{0, 0, 0, 0},
	}
	expected := 3
	result := manhattan.ManhattanDistance(matrix)

	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

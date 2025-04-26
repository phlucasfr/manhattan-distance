package manhattan

import "testing"

func BenchmarkManhattanDistanceWorstCase(b *testing.B) {
	const size = 100
	matrix := make([][]int, size)

	for i := range matrix {
		matrix[i] = make([]int, size)
		if i == size-2 {
			matrix[i][size-2] = 1
		}
		if i == size-1 {
			matrix[i][size-1] = 1
		}
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		ManhattanDistance(matrix)
	}
}

func BenchmarkManhattanDistanceAverageCase(b *testing.B) {
	matrix := [][]int{
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 1, 0},
		{0, 0, 0, 0, 1},
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		ManhattanDistance(matrix)
	}
}

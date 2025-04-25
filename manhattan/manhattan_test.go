package manhattan

import "testing"

func TestManhattanDistance(t *testing.T) {
	tests := []struct {
		name     string
		matrix   [][]int
		expected int
	}{
		{
			name: "Caso base do desafio",
			matrix: [][]int{
				{0, 0, 0, 0},
				{0, 1, 0, 0},
				{0, 0, 0, 1},
				{0, 0, 0, 0},
			},
			expected: 3,
		},
		{
			name:     "Matriz 1x2 com 1s adjacentes",
			matrix:   [][]int{{1, 1}},
			expected: 1,
		},
		{
			name: "Pontos em diagonal extrema",
			matrix: [][]int{
				{1, 0, 0},
				{0, 0, 0},
				{0, 0, 1},
			},
			expected: 4,
		},
		{
			name: "Segundo 1 na primeira posição",
			matrix: [][]int{
				{1, 0},
				{1, 0},
			},
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ManhattanDistance(tt.matrix)
			if result != tt.expected {
				t.Errorf(
					"Falha no teste: %s\nEsperado: %d\nObtido: %d\nMatriz: %v",
					tt.name,
					tt.expected,
					result,
					tt.matrix,
				)
			}
		})
	}
}

func TestEdgeCases(t *testing.T) {
	t.Run("Matriz mínima válida", func(t *testing.T) {
		matrix := [][]int{{1}, {1}}
		if got := ManhattanDistance(matrix); got != 1 {
			t.Errorf("Esperado 1, obtido %d. Matriz: %v", got, matrix)
		}
	})

	t.Run("Matriz esparsa grande", func(t *testing.T) {
		matrix := make([][]int, 100)
		for i := range matrix {
			matrix[i] = make([]int, 100)
		}
		matrix[0][0] = 1
		matrix[99][99] = 1
		if got := ManhattanDistance(matrix); got != 198 {
			t.Errorf("Esperado 198, obtido %d", got)
		}
	})
}

func TestAbs(t *testing.T) {
	t.Run("Valor positivo", func(t *testing.T) {
		if got := abs(5); got != 5 {
			t.Errorf("abs(5) = %d, esperado 5", got)
		}
	})

	t.Run("Valor negativo", func(t *testing.T) {
		if got := abs(-3); got != 3 {
			t.Errorf("abs(-3) = %d, esperado 3", got)
		}
	})

	t.Run("Zero", func(t *testing.T) {
		if got := abs(0); got != 0 {
			t.Errorf("abs(0) = %d, esperado 0", got)
		}
	})

	t.Run("Valor extremo negativo", func(t *testing.T) {
		if got := abs(-1_000_000); got != 1_000_000 {
			t.Errorf("abs(-1_000_000) = %d, esperado 1_000_000", got)
		}
	})
}

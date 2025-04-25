package manhattan

type Point struct {
	row, col int
}

// ManhattanDistance calcula a distância de Manhattan entre os dois únicos pontos com valor 1 em uma matriz bidimensional.
//
// A função assume que:
//   - A matriz é retangular (todas as linhas têm o mesmo comprimento)
//   - Existem exatamente dois elementos com valor 1 na matriz
//   - Os demais elementos são 0
//
// Parâmetros:
//   - matrix: Matriz de inteiros no formato [linhas][colunas], obrigatoriamente com dois elementos 1
//
// Retorna:
//   - Distância de Manhattan entre os dois pontos de valor 1, calculada pela fórmula:
//     |x1 - x2| + |y1 - y2|, onde (x,y) são as coordenadas dos pontos
func ManhattanDistance(matrix [][]int) int {
	var points [2]Point
	count := 0

	for i, row := range matrix {
		for j, val := range row {
			if val == 1 {
				points[count].row, points[count].col = i, j
				count++
				if count == 2 {
					break
				}
			}
		}
		if count == 2 {
			break
		}

	}

	rowDiff := points[0].row - points[1].row
	colDiff := points[0].col - points[1].col

	return abs(rowDiff) + abs(colDiff)
}

// abs retorna o valor absoluto de um número inteiro.
//
// Para números negativos (n < 0), retorna -n (inverso aditivo).
// Para zero ou positivos (n >= 0), retorna n inalterado.
//
// Parâmetros:
//   - number: Inteiro a ser processado
//
// Retorna:
//   - Sempre um valor não negativo: |n|
func abs(number int) int {
	if number < 0 {
		return -number
	}
	return number
}

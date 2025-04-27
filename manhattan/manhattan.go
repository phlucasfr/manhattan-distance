package manhattan

import (
	"github.com/phlucasfr/manhattan-distance/models"
	"github.com/phlucasfr/manhattan-distance/utils"
)

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
	var points [2]models.Points
	count := 0

	for i, row := range matrix {
		for j, val := range row {
			if val == 1 {
				points[count].Row, points[count].Col = i, j
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

	rowDiff := points[0].Row - points[1].Row
	colDiff := points[0].Col - points[1].Col

	return utils.Abs(rowDiff) + utils.Abs(colDiff)
}

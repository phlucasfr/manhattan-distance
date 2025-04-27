package utils

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
func Abs(number int) int {
	if number < 0 {
		return -number
	}
	return number
}

package utils

import "testing"

func TestAbs(t *testing.T) {
	t.Run("Valor positivo", func(t *testing.T) {
		if got := Abs(5); got != 5 {
			t.Errorf("abs(5) = %d, esperado 5", got)
		}
	})

	t.Run("Valor negativo", func(t *testing.T) {
		if got := Abs(-3); got != 3 {
			t.Errorf("abs(-3) = %d, esperado 3", got)
		}
	})

	t.Run("Zero", func(t *testing.T) {
		if got := Abs(0); got != 0 {
			t.Errorf("abs(0) = %d, esperado 0", got)
		}
	})

	t.Run("Valor extremo negativo", func(t *testing.T) {
		if got := Abs(-1_000_000); got != 1_000_000 {
			t.Errorf("abs(-1_000_000) = %d, esperado 1_000_000", got)
		}
	})
}

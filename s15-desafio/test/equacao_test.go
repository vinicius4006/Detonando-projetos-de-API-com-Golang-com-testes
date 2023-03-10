package test

import (
	"segundo-grau/entity"
	"testing"
)

func TestEquacaoSegundoGrau(t *testing.T) {
	t.Run("Calcular raízes", func(t *testing.T) {
		eq := entity.EquacaoSegundoGrau{A: 1, B: 3, C: 2}
		raiz1, raiz2 := eq.Calcular()
		esperadoRaiz1 := -1.0
		esperadoRaiz2 := -2.0

		if raiz1 != esperadoRaiz1 && raiz2 != esperadoRaiz2 {
			t.Errorf("result x1 %v expected x1 %v\nresult x2 %v expected x2 %v", raiz1, esperadoRaiz1, raiz2, esperadoRaiz2)
		}

	})
}

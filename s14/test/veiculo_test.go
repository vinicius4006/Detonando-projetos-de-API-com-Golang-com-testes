package test

import (
	"testing"

	"github.com/vinicius4006/detonando-projeto-api-golang/entity"
)

func TestVehicle(t *testing.T) {
	verificarVeiculoLigar := func(veiculo entity.Veiculo) {

		resultado := veiculo.Ligar()
		veiculoLigado := true
		if resultado != veiculoLigado {
			t.Errorf("resultado %v, esperado veiculoLigado %v", resultado, veiculoLigado)
		}
	}

	verificarVeiculoDesligar := func(veiculo entity.Veiculo) {
		resultado := veiculo.Desligar()
		veiculoDesligado := true
		if resultado != veiculoDesligado {
			t.Errorf("resultado %v, veiculoDesligado %v", resultado, veiculoDesligado)
		}
	}

	verificarVeiculoMover := func(veiculo entity.Veiculo) {
		resultado := veiculo.Mover(10, 50)
		veiculoMove := true

		if resultado != veiculoMove {
			t.Errorf("resultado %v, veiculoDesligado %v", resultado, veiculoMove)
		}

	}

	verificarVeiculoParar := func(veiculo entity.Veiculo) {
		veiculo.Parar()

	}
	t.Run("Carro", func(t *testing.T) {
		carro := entity.Carro{}
		t.Run("Ligar", func(t *testing.T) {

			verificarVeiculoLigar(carro)
		})

		t.Run("Desligar", func(t *testing.T) {
			verificarVeiculoDesligar(carro)
		})

		t.Run("Mover", func(t *testing.T) {
			verificarVeiculoMover(carro)
		})

		t.Run("Parar", func(t *testing.T) {
			verificarVeiculoParar(carro)
		})

	})

}

package main

import "github.com/vinicius4006/detonando-projeto-api-golang/entity"

func main() {
	carro1 := entity.Carro{}
	carro1.Ligar()
	carro1.Desligar()
	carro1.Mover(10, 50)
	carro1.Parar()
}

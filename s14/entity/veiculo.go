package entity

import (
	"fmt"
	"log"
)

type Veiculo interface {
	Ligar() bool
	Desligar() bool
	Mover(direcao, velocidade int) bool
	Parar()
}

type Carro struct {
}

func (c Carro) Ligar() bool {
	fmt.Println("Carro ligado...")
	return true
}

func (c Carro) Desligar() bool {
	fmt.Println("Carro desligado...")
	return true
}

func (c Carro) Mover(direcao, velocidade int) bool {
	fmt.Println("Carro movendo...")
	return true
}

func (c Carro) Parar() {
	_, err := fmt.Println("Parando...")
	if err != nil {
		log.Panicln(err)
	}
}

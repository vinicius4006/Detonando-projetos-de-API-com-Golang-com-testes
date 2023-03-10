package entity

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"sync"
)

type Equacao interface {
	Calcular() float64
}

type EquacaoSegundoGrau struct {
	A, B, C int
}

func (e EquacaoSegundoGrau) Calcular() (float64, float64) {
	var wg sync.WaitGroup
	var x1 float64
	var x2 float64
	wg.Add(2)
	a := float64(e.A)
	b := float64(e.B)
	c := float64(e.C)
	go func() {

		x1 = (-1*b + math.Sqrt(math.Pow(b, 2)-(4*a*c))) / 2 * a
		criarLog(x1)
		wg.Done()
	}()
	go func() {
		x2 = (-1*b - math.Sqrt(math.Pow(b, 2)-(4*a*c))) / 2 * a
		criarLog(x2)
		wg.Done()
	}()

	wg.Wait()
	if math.IsNaN(x1) || math.IsNaN(x2) {
		log.Fatalln("Raiz Negativa")
	}

	return x1, x2
}

func criarLog(v float64) {
	msg := fmt.Sprint(v)
	stringArr := []byte(msg)
	err := ioutil.WriteFile(fmt.Sprintf("logs/calculo-%v.txt", &msg), stringArr, 0644)
	if err != nil {
		log.Println(err)
	}
}

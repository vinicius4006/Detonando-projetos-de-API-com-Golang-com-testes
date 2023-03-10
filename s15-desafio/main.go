package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"segundo-grau/entity"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("(A)x^2 + (B)x + (C)\n")
	readerA := bufio.NewReader(os.Stdin)
	fmt.Print("A: ")
	a, erroa := readerA.ReadString('\n')
	check(erroa)
	A, erroA := strconv.Atoi(fmt.Sprint(strings.TrimSpace(a)))
	check(erroA)
	//
	readerB := bufio.NewReader(os.Stdin)
	fmt.Print("B: ")
	b, errob := readerB.ReadString('\n')
	check(errob)
	B, erroB := strconv.Atoi(fmt.Sprint(strings.TrimSpace(b)))
	check(erroB)
	//
	readerC := bufio.NewReader(os.Stdin)
	fmt.Print("C: ")
	c, erroc := readerC.ReadString('\n')
	check(erroc)
	C, erroC := strconv.Atoi(fmt.Sprint(strings.TrimSpace(c)))
	check(erroC)
	//
	fmt.Printf("Calculando...\n")
	eqS := entity.EquacaoSegundoGrau{A: A, B: B, C: C}
	x1, x2 := eqS.Calcular()
	fmt.Printf("Raíz 1: %v\nRaíz 2: %v", x1, x2)

}

func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

package main

import "fmt"

type Solution struct{}

func main() {
	vetores := []int{2, 7, 11, 15, 4, 9, 4, 5, 78, 21, 65, 87, 98, 21, 5874, 5, 61, 32, 32, 32}
	somaValor := 14

	minhaSolucaoSoma(vetores, somaValor)

}

func minhaSolucaoSoma(vetores []int, alvo int) {

	mapVetores := make(map[int]int)

	for idx, value1 := range vetores {
		mapVetores[value1] = idx
	}
	for idx2, value2 := range vetores {
		complemento := alvo - value2
		_, ok := mapVetores[complemento]
		if ok && mapVetores[complemento] != idx2 {
			fmt.Printf("%v %v", value2, complemento)
			break
		}
	}

}

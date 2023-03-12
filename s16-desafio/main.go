package main

import "fmt"

type Solution struct{}

func (s Solution) twoSums(nums []int, target int) []int {
	dictValues := make(map[int]int)

	for ix, val := range nums {
		dictValues[val] = ix
	}

	for ix, val := range nums {
		complement := target - val
		if _, ok := dictValues[complement]; ok && dictValues[complement] != ix {
			return []int{ix, dictValues[complement]}
		}
	}
	return nil
}

func main() {
	vetores := []int{2, 7, 11, 15, 4, 9, 4, 5, 78, 21, 65, 87, 98, 21, 5874, 5, 61, 32, 32, 32}
	somaValor := 67
	// minha solucacao
	minhaSolucaoSoma(vetores, somaValor)

	//solucao professor
	var s Solution
	arr := s.twoSums(vetores, somaValor)
	fmt.Printf("%v %v", arr[0], arr[1])

}

func minhaSolucaoSoma(vetores []int, alvo int) {
	stopFor := false
	inicialPositionVetor := 0
	position := 0
	newArray := []int{}
	for {

		var soma int
		if stopFor || inicialPositionVetor >= len(vetores) {
			break
		}
		if position < len(vetores) {
			if position == inicialPositionVetor && inicialPositionVetor < len(vetores)-1 {
				position++
			}
			soma = vetores[inicialPositionVetor] + vetores[position]
			if soma == alvo {
				newArray = append(newArray, inicialPositionVetor, position)
				stopFor = true
			}
			position++
		} else {
			inicialPositionVetor++
			position = 0
		}

	}
	fmt.Printf("%v", newArray)
}

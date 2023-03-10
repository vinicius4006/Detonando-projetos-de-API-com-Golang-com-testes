package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func checked(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Your name: ")
	text, _ := reader.ReadString('\n')
	fmt.Printf("Hello, %v how are you today\n", text)

	newName := fmt.Sprintf("%v", strings.TrimSpace(text))
	fmt.Printf("Now, %v, without the extra line\n", newName)

	//numeros inteiros
	fmt.Print("Now, type an integer: ")
	nints, _ := reader.ReadString('\n')
	nint, errInt := strconv.ParseInt(strings.TrimSpace(nints), 10, 64)
	checked(errInt)
	fmt.Printf("You typed: %d\n", nint)

	//numeros reais
	fmt.Print("Now, type a float: ")
	nfs, _ := reader.ReadString('\n')
	nf, errf := strconv.ParseFloat(strings.TrimSpace(nfs), 64)
	checked(errf)
	fmt.Printf("You typed: %f\n", nf)

	//internacionalizacao
	p := message.NewPrinter(language.BrazilianPortuguese)
	p.Printf("%f", nf)

	//Creating a text file
	stringArr := []byte("Bom dia.\nComa uma pêra!\nTenha um ótimo dia!\n")
	// Permission: -rw-r--r--
	err := ioutil.WriteFile("G:/Projetos/Go/Detonando projetos de API com Golang/s15/arq.txt", stringArr, 0644)
	checked(err)

	//Lendo o arquivo
	data, err := ioutil.ReadFile("G:/Projetos/Go/Detonando projetos de API com Golang/s15/arq.txt")
	checked(err)
	fmt.Printf("\nType: %T\n", data)
	textContent := string(data)
	fmt.Println(textContent)

}

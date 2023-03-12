package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	chAnswer := make(chan string)

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("Type a term number or other character to finish: ")
		nints, _ := reader.ReadString('\n')
		nint, erInt := strconv.ParseInt(strings.TrimSpace(nints), 10, 64)

		if erInt != nil {
			break
		}
		go FibonacciLoop(int(nint), chAnswer)
		select {
		case answer := <-chAnswer:
			fmt.Println("Got an answer: ", answer)
		default:
			fmt.Println("Waiting...")
		}
	}

}

func FibonacciLoop(n int, aChannel chan string) {
	f := make([]int, n+1, n+2)
	if n < 2 {
		f = f[0:2]
	}
	f[0] = 0
	f[1] = 1
	for i := 2; i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
	}

	aChannel <- fmt.Sprintf("The term %d is %d\n", n, f[n])
}

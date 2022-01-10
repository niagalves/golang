// é a unica estrutura de repetição no go, mas existem diversas formas de uso
package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1

	for i <= 10 {
		fmt.Println(i)
		i++
	}

	for i := 0; i <= 20; i++ {
		fmt.Println("Número", i)
	}

	// laço infinito
	for {
		fmt.Println("Para sempre..")
		time.Sleep(time.Second)
	}
}

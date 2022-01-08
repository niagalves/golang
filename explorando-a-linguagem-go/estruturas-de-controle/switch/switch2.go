package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()

	// switch true, vai procurar o primeiro case que seja verdadeiro
	switch {
	case t.Hour() < 12:
		fmt.Println("Bom dia!")
	case t.Hour() < 18:
		fmt.Println("Boa tarde!")
	default:
		fmt.Println("Boa noite!")
	}
}

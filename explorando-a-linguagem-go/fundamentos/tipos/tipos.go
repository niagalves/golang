package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// números inteiros
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é", reflect.TypeOf(32000))

	// sem sinal (só positivos)... uint8, uint16, uint32, uint64
	// o que varia é o tamanho

	var b byte = 255
	fmt.Println("o bite é", reflect.TypeOf(b))

	// com sinal.. int8, int16 int32, int64
	i1 := math.MaxInt64
	fmt.Println("o valor máximo do int é", i1)
	fmt.Println("o tipo de i1 é", reflect.TypeOf(i1))

	var i2 rune = 'a' // representa um mapeamento da tabela unicode (int32)
	fmt.Println("o rune é", reflect.TypeOf(i2))
	fmt.Println(i2)

	// números reais
	var x = float64(49.99)
	fmt.Println("o tipo x é", reflect.TypeOf(x))

	// boolean
	var bo = false
	fmt.Println("o tipo é", reflect.TypeOf(bo))
	fmt.Println(bo)

	// string
	s1 := "olá mundo"
	fmt.Println("o tipo é", reflect.TypeOf(s1))
	fmt.Println(s1)

	// string com multiplas linhas
	s2 := `
		olá
		bem vindo
		ao mundo
		feliz
	`
	fmt.Println("o tipo é", reflect.TypeOf(s2))
	fmt.Println(s2)

	// char não existe em go.
	c := "a"
	fmt.Println("o tipo é", reflect.TypeOf(c))
	fmt.Println(c)
}

package main

// quando preciso ter o import, mas sem utilizar da para por um _ na frente:
// import (
// 	"fmt"
// 	_"math"
// )

// podemos renomear os pacotes utilizando assim:
// import (
// 	"fmt"
// 	m "math"
// )

import (
	"fmt"
	"math"
)

func main() {
	const PI float64 = 3.1415 //float64 é o padrão
	var raio = 3.2

	// forma reduzida (essa é a forma de preferencia)
	area := PI * math.Pow(raio, 2)

	fmt.Println("Área da circunferência é", area)

	// outras formas
	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)
}

// golang é fortimente tipada

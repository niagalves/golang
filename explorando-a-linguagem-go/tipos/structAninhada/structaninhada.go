package main

import "fmt"

type item struct {
	produtoID int
	qtde      int
	preco     float64
}

type pedido struct {
	userID int
	itens  []item
}

func (p pedido) valorTotal() float64 {
	total := 0.0

	for _, item := range p.itens {
		total += item.preco * float64(item.qtde)
	}

	return total
}

func main() {
	pedido := pedido{
		userID: 1,
		itens: []item{
			{1, 2, 3.50},
			{2, 1, 33.50},
			{11, 199, 2.50},
		},
	}

	fmt.Printf("O valor total do pedido é R$ %.2f", pedido.valorTotal())
}

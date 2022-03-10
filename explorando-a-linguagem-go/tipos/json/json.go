package main

import (
	"encoding/json"
	"fmt"
)

type produto struct {
	ID    int      `json:"id`
	Nome  string   `json:"nome"`
	preco float64  `json:"preco"`
	Tags  []string `json:"tags`
}

func main() {
	p1 := produto{1, "notebook", 1899.90, []string{"Promoção", "Eletrônico"}}
	p1Json, _ := json.Marshal(p1)
	fmt.Println(string(p1Json))

	// json para struct
	var p2 produto
	jsonString := `{"id": 2, "nome": "caneta", "preco": 8.90, "tags": ["papelaria", "teste"]}`
	json.Unmarshal([]byte(jsonString), &p2)

	fmt.Println(p2.Tags[1])
}

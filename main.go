package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	contaDoLucas := ContaCorrente{titular: "Lucas",
		saldo: 125.5,
	}
	contaDaBruna := ContaCorrente{"Bruna", 588, 123457, 125.5}

	fmt.Println(contaDoLucas)
	fmt.Println(contaDaBruna)
}

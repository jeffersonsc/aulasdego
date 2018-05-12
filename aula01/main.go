package main

import "fmt"

func main() {
	var endereco string
	var idade int
	var saldo float64

	fmt.Printf("Endereco é %s \r\n", endereco)
	fmt.Printf("Idade é %d \r\n", idade)
	fmt.Printf("Saldo é %v \r\n", saldo)

	a := "meu texto" // atribui uma string
	b := 10          // atribui um int
	c := 10.5        // atribui um float
	d := true        // atriui um booleano

	fmt.Printf("A: %s, B: %d, C: %v, D: %v \r\n", a, b, c, d)

	var nome, sobrenome string
	var corrente, poupanca float64

	nome = "Jeff"
	sobrenome = "Cunha"

	fmt.Printf("Nome completo: %s %s \r\n", nome, sobrenome)
	fmt.Printf("Conta corrente: %v, Poupança: %v \r\n", corrente, poupanca)

	var chamaFuncao = func() {
		fmt.Print("Chamando funcao\r\n")
	}

	chamaFuncao()
}

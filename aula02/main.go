package main

import (
	"fmt"

	"github.com/jeffersonsc/cursogo/aula02/fibonacci"

	"github.com/jeffersonsc/cursogo/aula02/fatorial"
)

func main() {
	fmt.Print("Hello World \r\n")

	resultado := soma(1, 3)
	fmt.Printf("Resultado da soma: %d \r\n", resultado)

	minhaIdade, meuNome := myFuncionVar()
	fmt.Printf("Meu nome e idade é: %s, %d\r\n", meuNome, minhaIdade)

	meuNome, minhaIdade = MyFunctionVar()
	fmt.Printf("Meu nome e idade é: %s, %d\r\n", meuNome, minhaIdade)

	var chamaFuncao = func() {
		fmt.Print("Chamando função. \r\n")
	}

	chamaFuncao()

	myFatorial := fatorial.Fatorial(4)
	fmt.Printf("Resultado do fatorial: %d \r\n", myFatorial)

	myFib := fibonacci.Fibonacci(6)
	fmt.Printf("Resultado do fibonacci: %d \r\n", myFib)

}

func soma(a, b int) int {
	return a + b
}

func privada() {} // apenas os arquivos do mesmo pacote veem
// Publica precisasmos comentar para documentar essa exportação de função para fora.
func Publica() {} // todos que instanciarem esse pacote veem essa função

func myFuncion() (int, string) {
	return 100, "Minha função"
}

func myFuncionVar() (int, string) {
	nome := "Jeff"
	idade := 23

	return idade, nome
}

// MyFunctionVar mesmo nome porém publica
func MyFunctionVar() (nome string, idade int) {
	nome = "Jeff"
	idade = 23

	return
}

## Pacote (packages) e Funções ##

- No go pacotes são definidos pelo nome da pasta exeto o pacote main que fica na raiz do projeto onde 
é inicializado o projeto e executado. 
- Não é abitual o go executar o arquivos que não estejam com o pacote main

```
package main

func main() {
  fmt.Print("Hello World \r\n")
}
```

No exemplo acima estamos falando para o aquivo que o ele é o main e chamamos do pacote "fmt" a função "Print"

### Funções ###

- Para definir uma função usamos a palavra reservada "func" seguida do nome da função, argumentos e seu retorno o go pode retornar mais de um valor

```
  func soma(a,b int) int {
    return a + b
  }

```

- Para falar que uma função é privada ou pública definimos pelo seu nome se o começo for maiúsculo é publica se for minúsculo.

exemplo:

```
  func privada(){} // apenas os arquivos do mesmo pacote veem
  func Publica(){} // todos que instanciarem esse pacote veem essa função
```

### Retorno de funções ###

- Para definir o retorno de uma fucção usamos a palavra reservado chamada de "return" em seguida passando os valores de retorno

```
func myFuncion() (int,string) {
  return 100, "Minha função"
}

func myFuncionVar() (int,string) {
  nome := "Jeff"
  idade := 23

  return idade, nome
}
```

### Nomeando retorna da função ###

- para facilitar nosso trabalho de de ficar declarando variáveis para passar no retorno o go facilita

```
func MyFunctionNew() (nome string, idade int) {
  nome = "Jeff"
  idade = 23

  return
}
```

o que acontece em cima é que o go já declara as variáveis de retorno de função e apenas passamos o valor e damos um return e ele já sabe o que retornar por que as variáveis de retorno.


### Funções anonimas ###

- funções animas são aquelas que não tem seus nomes declarados explicitamente mas implicitamente está lá vamos ao exemplo para ficar claro nesse exemplo passamos uma função para uma variável.

```
var chamaFuncao = func(){
  fmt.Print("Chamando função. \r\n")
}

chamaFuncao()
```

Aciama falamos que essa variável tem o papel de guargar essa função e uma vez chamada ela a função é acionada.

## argumentos de função ##

- podemos passar os tipo primitivos, passar funções e estruturas de dados e ponteiros de dados.

```
  func soma(a,b int) int {
    return a + b
  }

  result := soma(1,2)

```

- podemos passar os ponteiros

```
  func soma(res *int, a,b int) {
    res = a + b
  }

  var result int
  soma(&result, 4,2)

  // Print de result

  var idade = 3
  fmt.Printf("minha idade: %d \r\n", idade)
  somai := somaIdade(&idade)
  fmt.Printf("minha idade++: %d \r\n", somai)
  fmt.Printf("minha idade: %d \r\n", idade)

  func somaIdade(idade *int) int {
    return idade++
  }

```

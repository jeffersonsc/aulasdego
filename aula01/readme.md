## Brincando com variáveis ##

### Tipo primitivos ###

- string textos UFT-8
- int inteiros
- float64 decimais
- bool boleanos true e false
- rune para usar com caracteres especiais

### Atribuição de variáveis ###
- para atribuir uma variávei usamos a sequencia do seguida de var nome tipo ao contrário de muitas linguagens o go define do tipo apos o nome.
```
var nome string
var idade int
var saldo float64
```
- para atribuir um valor a uma variável sem precisar declarar um var usamos os caractarer dois pontos ':' junto com o igual

```
  a := "meu texto" // atribui uma string
  b := 10 // atribui um int
  c := 10.5 // atribui um float
  d := true // atriui um booleano
  e := User{} // atribuiu uma estrutura
```

- podemos atribuir para varias variáveis o mesmo tipo separando elas apenas por ',' vírgula

```
  var nome, sobrenome string
  var corrente, popanca float64
  
```

- Variáveis em go pode aceitar diversos tipos de dados desde que estejam de acordo com o permitido. 

```
  var chamaFuncao = func(){
    fmt.Print("Chamando funcao\r\n")
  }()

  chamaFuncao
```

- Para diferenciar variáveis publicas de privadas usamos o nome dela para definir, maiúscula é pública minúcula é privada
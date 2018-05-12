package main

import (
	"encoding/json"
	"fmt"
)

// User estrutura publica publica
type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`

	Address Address `json:"adressess"`

	Apelidos []Nickname `json:"apelidos"`
}

// user estrutura privada
type user struct {
	name string
	age  int
}

// Address estrutura de endereço
type Address struct {
	Name string
	Num  int
}

// Nickname apelidos
type Nickname struct {
	Name string
}

var (
	user1 User
	user2 user
)

func main() {
	user1.Name = "Jeff"
	user1.Age = 23

	user1.Address.Name = "Salvador penteado"
	user1.Address.Num = 366

	user1.Apelidos = []Nickname{
		Nickname{Name: "Jeff"},
	}

	resutado, err := json.Marshal(user1)
	if err != nil {
		fmt.Printf("Error ao parsear json. ERROR: %s", err.Error())
		return
	}

	fmt.Print(string(resutado))

	user2.name = "Jeff"
	user2.age = 23

}

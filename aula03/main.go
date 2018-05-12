package main

import (
	"fmt"
)

func main() {
	a := 1
	b := 2

	c := a + b
	d := a * b
	e := a / b
	f := b - a

	fmt.Printf("c = %d, d = %d, e = %d, f = %d\r\n", c, d, e, f)

	var slice []int
	var array [3]int

	slice = []int{1, 2, 3, 4, 5, 6}

	slice = append(slice, 7)

	array[0] = 1
	array[1] = 2
	array[2] = 3

	fmt.Printf("slice: %v. array: %v \r\n", slice, array)

}

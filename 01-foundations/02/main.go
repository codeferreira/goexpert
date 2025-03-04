package main

import "fmt"

type Numeric interface {
	int | float64
}

func Soma[Type Numeric](a, b Type) Type {
	return a + b
}

func main() {
	a := Soma(1, 2)
	b := Soma(1.0, 2.0)

	fmt.Println(a)
	fmt.Println(b)
}

package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	defer fmt.Println("Hello, World! 2")
	fmt.Println("Hello, World! 3")
}

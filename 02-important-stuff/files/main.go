package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("./test.txt")
	if err != nil {
		panic(err)
	}

	size, err := file.WriteString("Hello World\n")
	if err != nil {
		panic(err)
	}

	fmt.Printf("File created with size %d\n", size)

	size2, err := file.Write([]byte("Second line\n"))
	if err != nil {
		panic(err)
	}

	fmt.Printf("File updated with size %d\n", size2)
	file.Close()

	f, err := os.ReadFile("test.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(f))

	f2, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(f2)
	buffer := make([]byte, 10)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}

	removeErr := os.Remove("test.txt")
	if removeErr != nil {
		panic(removeErr)
	}
}

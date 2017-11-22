package main

import "fmt"

const p string = "defining constant"

func main() {
	const q = 42
	fmt.Println("p - ", p)
	fmt.Println("q - ", q)
	// Constant is a simple unchanging value
}

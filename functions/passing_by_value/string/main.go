package main

import "fmt"

func main() {

	name := "Haps"
	fmt.Println(name) // Todd

	changeMe(name)

	fmt.Println(name) // Todd
}

func changeMe(z string) {
	fmt.Println(z) // Todd
	z = "Rocky"
	fmt.Println(z) // Rocky
}
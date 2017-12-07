package main

import "fmt"

func main() {
	greet("Haps")
	greet("Me")
}

func greet(name string) {
	fmt.Println("Hello", name)
}

// greet is declared with a parameter
// when calling greet an argument is passed
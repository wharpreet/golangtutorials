package main

import "fmt"

func main() {
	greet("Haps", "Me")
}

func greet(fname string, sname string) {
	fmt.Println("Hello", fname, "and", sname)
}

// greet is declared with a parameter
// when calling greet an argument is passed

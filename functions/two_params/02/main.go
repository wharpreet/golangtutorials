package main

import "fmt"

func main() {
	greet("Haps", "Me")
}

func greet(fname, sname string) {
	fmt.Println("Hello", fname, "and", sname)
}

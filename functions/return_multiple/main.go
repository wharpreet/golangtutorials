package main

import "fmt"

func main() {
	fmt.Println(greet("Haps", "Me"))
}

func greet(fname, sname string) (string, string) {
	// Sprint (String Print) prints to a string and not a standard out
	return fmt.Sprint("Hello ", fname, " and ", sname), fmt.Sprint("Hello ", sname, " and ", fname)
}
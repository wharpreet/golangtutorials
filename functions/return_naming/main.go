package main

import "fmt"

func main() {
	fmt.Println(greet("Haps", "Me"))
}

func greet(fname, sname string) (s string) {
	// Sprint (String Print) prints to a string and not a standard out
	s = fmt.Sprint("Hello ", fname, " and ", sname)
	return
}

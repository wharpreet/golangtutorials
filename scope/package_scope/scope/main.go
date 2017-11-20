package main

import "fmt"

//order of declaration and execution matters
//package level scope
var x int = 42

func main() {
	fmt.Println(x)
	foo()

	//block level scope
	y := 17
	fmt.Println(y)
}
func foo() {
	fmt.Println(x)
}

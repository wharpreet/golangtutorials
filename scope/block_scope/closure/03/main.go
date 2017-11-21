package main

import "fmt"

func main() {
	x := 0
	// anonymous function: when no name is assigned to function
	//func expression: when a fucntion is assigned to a variable
	increment:= func() int {
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())
}

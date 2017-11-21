package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)
	{
		fmt.Println(x)
		y := "Printing from inner closure"
		fmt.Println(y)
	}
	// If we try to print "y" here outside the inner closure, it wont work
	// fmt.Println(y)
}

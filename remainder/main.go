package main

import "fmt"

func main() {
	x := 14 % 3
	fmt.Println("Remainder of 14 /3 is",x)

	if x == 1 {
		fmt.Println("Odd")
	} else {
		fmt.Println("Even")
	}

	for i := 1; i < 100; i++ {
		if i % 2 == 1 {
			fmt.Println("Odd")
		} else {
			fmt.Println("Even")
		}
	}
}

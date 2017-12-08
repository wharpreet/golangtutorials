package main

import "fmt"

func main() {
	fmt.Println((true && false) || (false && true) || !(false && false))
	if (true && false) || (false && true) || !(false && false) {
		fmt.Println("It's true")
	}
}

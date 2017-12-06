package main

import "fmt"

func main() {
	// Only "!" "false" will run every time in this file, as that is returns true
	if !true {
		fmt.Println("It's not true")
	}
	if !false {
		fmt.Println("It's not flase")
	}
}

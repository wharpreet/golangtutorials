package main

import "fmt"

func main() {
	b := true
	// Food variable in being initialised in "if" statement
	// This keeps the skope of the variable very tight and limited only to the if statement
	if food := "Chocolate"; b {
		fmt.Println(food)
	}
}

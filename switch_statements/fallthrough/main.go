package main

import "fmt"

func main() {
	switch "Best" {
	case "Test":
		fmt.Println("Wassup Test")
	case "Best":
		fmt.Println("Wassup Best")
		// fallthrough assumes next one is also true
		fallthrough
	case "abc":
		fmt.Println("Wassup abc")
		fallthrough
	case "xyz":
		fmt.Println("Wassup xyz")
	case "Appu":
		fmt.Println("Wassup Appu")
	default:
		fmt.Println("Have you no friends?")
	}
}

/*
no default fallthough
fallthough is optional
-- you can specify fallthough by explicitly stating it
-- break isn't needed like in other languages
*/

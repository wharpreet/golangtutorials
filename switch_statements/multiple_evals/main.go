package main

import "fmt"

func main() {
	switch "Best" {
	case "Test", "Best":
		fmt.Println("Wassup Test, or, Best")
	case "abc", "xyz":
		fmt.Println("Wassup abc, or, xyz")
	case "Appu", "Aps":
		fmt.Println("Wassup Appu, or, Aps")
	}
}

package main

import "fmt"

func main() {
	n := average(43, 56, 87, 12, 45, 57)
	fmt.Println(n)
}

// Veriadic is the last argument of any function where variable type is prefixed with "..."
// This variable can take zero to any number of arguments
func average(sf ...float64) float64 {
//	fmt.Println(sf)
//	fmt.Printf("%T \n", sf)
	total := 0.0
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}

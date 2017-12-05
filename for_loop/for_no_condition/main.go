package main

import "fmt"

func main() {
	i := 0
	// This loop will keep going on and on as there is no break defined
	for {
		fmt.Println(i)
		i++
	}
}

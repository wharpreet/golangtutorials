package main

import "fmt"

func main() {
	i := 0
	// This loop will break after 10 due to break specified in IF condition
	for {
		fmt.Println(i)
		if i >= 10 {
			break
		}
		i++
	}
}

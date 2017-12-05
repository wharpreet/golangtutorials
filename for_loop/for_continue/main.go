package main

import "fmt"

func main() {
	i := 0
	for {
		i++
		// This loop will print only odd number, as print statement is after first IF that check for remainder when i divided by 2
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
		// And second IF statement will break loop once it finds an odd number greater than 50
		if i >= 50 {
			break
		}
	}
}

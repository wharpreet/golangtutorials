package main

import "fmt"

// here callback is the name of the function
func visit(numbers []int, callback func(int)) {
	for _, n := range numbers {
		callback(n)
	}
}
func main() {
	visit([]int{1, 2, 3, 4}, func(n int) {
		fmt.Println(n)
	})
}

// callback: passing of function as an argument
// Go is simple, this should be avioded in go
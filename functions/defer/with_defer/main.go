package main

import "fmt"

func hello() {
	fmt.Print("Hello ")
}
func world() {
	fmt.Println("World")
}
func main() {
	// this defers the running of world function till the end of the main function
	defer world()
	hello()
}

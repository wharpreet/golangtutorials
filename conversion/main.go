package main

import "fmt"

func main() {
	// Shows byte for each english character
	fmt.Println([]byte("Hello"))
	// Shows bytes used by each chinese character - three bytes per character
	fmt.Println([]byte("世界"))
}

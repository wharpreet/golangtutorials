package main

import "fmt"

func main() {
	for i := 50; i <= 140; i++ {
		fmt.Println(i, "-", string(i), "-", []byte(string(i)))
	}
	// Single quote defines the rune
	foo := 'a'
	fmt.Println(foo)
	fmt.Printf("%T \n", foo)
	fmt.Printf("%T \n", rune(foo))

	// Now replacing singl quotes with double to see if that prints the string
	fo := "a"
	fmt.Println(fo)
	fmt.Printf("%T \n", fo)
}

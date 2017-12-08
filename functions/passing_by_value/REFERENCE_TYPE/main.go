package main

import "fmt"

func main() {
	m := make(map[string]int)
	changeMe(m)
	fmt.Println(m["Haps"]) // 34
}

func changeMe(z map[string]int) {
	z["Haps"] = 34
}
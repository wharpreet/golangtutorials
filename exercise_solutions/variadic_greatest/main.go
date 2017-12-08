package main
import "fmt"
func main() {
	greatest := max(4, 7, 10, 225, 120, 55, 240)
	fmt.Println(greatest)
}
func max(numbers ...int) int {
	var largest int
	for _, v := range(numbers) {
		if v > largest {
			largest = v
		}
	}
	return largest
}
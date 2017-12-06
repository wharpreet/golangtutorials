package main

import "fmt"

func main() {
	MyFriendsName := ""
	switch {
	case len(MyFriendsName) == 2:
		fmt.Println("Wassup my friend with name of length 2")
	case MyFriendsName == "Test":
		fmt.Println("Wassup Test")
	case MyFriendsName == "Best":
		fmt.Println("Wassup Best")
	case MyFriendsName == "Appu", MyFriendsName == "Aps":
		fmt.Println("Your name is either Appu or Aps")
	default:
		fmt.Println("Nothing matched, this is the default")
	}
}

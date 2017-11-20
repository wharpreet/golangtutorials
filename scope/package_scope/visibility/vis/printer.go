package vis

import "fmt"

func PrintVar() {
	fmt.Println(MyName)
	//Even though yourName starts with lower case, it is still visible in different file, as they both are part of same package
	fmt.Println(yourName)
}

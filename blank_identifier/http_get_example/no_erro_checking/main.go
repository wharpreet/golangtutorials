package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// Response is ignore with "_"
	res, _ := http.Get("http://www.google.com")
	page, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s", page)
}

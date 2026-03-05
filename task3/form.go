package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Content-Type: text/html")
	fmt.Println()

	err := (&http.Request{}).ParseForm()

	if err != nil {
		fmt.Println("Form reading error")
	}
}

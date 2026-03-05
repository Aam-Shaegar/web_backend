package main

import (
	"fmt"
	"net/http"
	"net/http/cgi"
)

func handler(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		fmt.Fprintln(w, "Form reading error")
		return
	}

	fmt.Fprintln(w, "<h1>CGI работает</h1>")
}

func main() {
	cgi.Serve(http.HandlerFunc(handler))
}

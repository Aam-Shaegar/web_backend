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

	name := r.FormValue("name")
	phone := r.FormValue("phone")
	email := r.FormValue("email")
	birthday := r.FormValue("birthdate")
	gender := r.FormValue("gender")
	bio := r.FormValue("bio")
	contract := r.FormValue("contract")
	languages := r.Form["languages[]"]
	fmt.Fprintln(w, "<h1>CGI работает</h1>")

	fmt.Fprintf(w, `
	<html>
	<head>
	<title>Form result</title>
	</head>
	<body>

	<h1>Submitted data</h1>

	Name: %s <br>
	Phone: %s <br>
	Email: %s <br>
	Birthdate: %s <br>
	Gender: %s <br>
	Bio: %s <br>
	Contract accepted: %s <br>

	<h3>Languages:</h3>
	<ul>
	`, name, phone, email, birthday, gender, bio, contract)

	for _, lang := range languages {
		fmt.Fprintf(w, "<li>%s</li>", lang)
	}

	fmt.Fprintf(w, `
	</ul>

	</body>
	</html>
	`)

}

func main() {
	cgi.Serve(http.HandlerFunc(handler))
}

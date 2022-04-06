package main

import (
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is The Home Page")
}

func About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is The About Page")
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println("Server are running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

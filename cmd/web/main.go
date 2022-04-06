package main

import (
	"fmt"
	"hotel-manage/package/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println("Server are running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

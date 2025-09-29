package main

import (
	"fmt"
	"io"
	"net/http"
)

func get(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Got you body")
	io.WriteString(w, "Got you budy!")
}
func main() {
	http.HandleFunc("/", get)
	http.ListenAndServe(":8080", nil)
	fmt.Println("Serving...")
}

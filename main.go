package main

import "net/http"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Handle the root URL
		// Write "Hello, World!" to the response
		w.Write([]byte("<h1>Hello, World!</h1>"))
	})
	// This is a simple Go program that prints "Hello, World!" to the console.
	http.ListenAndServe(":8080", nil)

}

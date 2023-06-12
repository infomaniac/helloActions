package main

import "net/http"

func main() {
	http.ListenAndServe(":8080", http.HandlerFunc(helloHandler))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!\n"))
}

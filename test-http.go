package main

import "net/http"

func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":3336", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World Web3"))
}

package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello")
}

func main() {
	http.HandleFunc("/", handlerFunc)
	fmt.Println("see server at port 3001")
	http.ListenAndServe(":3001", nil)
}

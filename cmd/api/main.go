package main

import "net/http"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /calculate", calculateGradeHandler)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err)
	}
}

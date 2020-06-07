package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/randomFloat", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, rand.Float64())
	})

	mux.HandleFunc("/randomInt", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, rand.Intn(100))
	})

	http.ListenAndServe(":8000", mux)
}

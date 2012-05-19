package main

import (
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("")))
	if err := http.ListenAndServe(":9090", nil); err != nil {
		panic(err)
	}
}

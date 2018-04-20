package main

import (
	"os"
	"io"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	name, err := os.Hostname()

    if err != nil {
        panic(err)
    }

	io.WriteString(w, "Hello " + name)
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":9090", nil)
}
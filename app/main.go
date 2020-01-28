package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})

	port, exists := os.LookupEnv("PORT")
	if !exists {
		port = "8088"
	}
	http.ListenAndServe(":" + port, nil)
}
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("RequestURI: %s ", r.RequestURI)
		n, err := fmt.Fprint(w, "Hello world!")
		if err != nil {
			fmt.Printf("Error found: %s\n", err)
		}
		fmt.Printf("- bytes written: %d\n", n)
	})

	http.ListenAndServe(":8083", nil)
}

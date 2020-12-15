package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "OK")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "4321"
	}
	fmt.Printf("HTTP server listening on http://localhost:%s\n", port)
	http.ListenAndServe(":"+port, nil)
}

package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := setRouting()

	fmt.Println("<---- Starting the server ---->")
	server := http.Server{
		Addr:    ":8080",
		Handler: wrapper(mux),
	}
	server.ListenAndServe()
}

func wrapper(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		h.ServeHTTP(w, r)
	})
}

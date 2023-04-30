package main

import (
	"fmt"
	"net/http"
)

// definition of a function in Func type
func Middleware(h http.Handler) http.Handler {
	// using http.HandlerFunc, you can create http.Handler from a function
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")
		fmt.Println("Authorization: ", authorization)

		h.ServeHTTP(w, r)
	})
}

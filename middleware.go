package main

import (
	"database/sql"
	"net/http"
)

// As it's difficult read the Middleware function in https://gqlgen.com/recipes/authentication/
// I'm summarizing types used in the function.
//
// *sql.DB       -> |Middleware| -> Func
//  http.Handler -> |   Func   | -> http.Handler

func Middleware(db *sql.DB) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return nil
	}
}

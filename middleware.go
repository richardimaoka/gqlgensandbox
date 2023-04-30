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

type Func func(http.Handler) http.Handler

func Middleware(db *sql.DB) Func {

	// definition of a function in Func type
	return func(next http.Handler) http.Handler {
		return nil
	}
}

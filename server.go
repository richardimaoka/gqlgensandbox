package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	firebase "firebase.google.com/go/v4"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/richardimaoka/gqlgensandbox/graph"
)

const defaultPort = "8080"

// receive http.Handler and return http.Handler
func Middleware(graphQLHandler http.Handler) http.Handler {
	// using http.HandlerFunc, you can create http.Handler from a function
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")
		fmt.Printf("Authorization: `%s`\n", authorization)

		if authorization != "correct token" {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
		} else {
			graphQLHandler.ServeHTTP(w, r)
		}
	})
}

func main() {
	/*app*/ _, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	if os.Getenv("AUTH_NEEDED") == "true" {
		http.Handle("/", Middleware(playground.Handler("GraphQL playground", "/query")))
		http.Handle("/query", Middleware(srv))
	} else {
		http.Handle("/", playground.Handler("GraphQL playground", "/query"))
		http.Handle("/query", srv)
	}

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

package main

import (
	"log"
	"net/http"
	"os"

	"github.com/conradwt/gqlgen_example"
	"github.com/conradwt/gqlgen_example/graph/generated"
	"github.com/vektah/gqlgen/handler"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(generated.NewExecutableSchema(generated.Config{Resolvers: &gqlgen_example.Resolver{}, Directives: &generated.DirectiveRoot{}})))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

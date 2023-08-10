package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/hwoodall30/sqlite-gql/database"
	datasources "github.com/hwoodall30/sqlite-gql/graph/dataSources"
	"github.com/hwoodall30/sqlite-gql/graph/generated"
	"github.com/hwoodall30/sqlite-gql/graph/resolvers"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db := &database.DataBase{}
	db.InitDatabase()

	dataSources := &datasources.DataSource{DB: db}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolvers.Resolver{DataSources: dataSources}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

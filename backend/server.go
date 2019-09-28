package main

import (
	"log"
	"net/http"
	"os"

	"github.com/graphql-todo/backend/graphql"

	"github.com/graphql-todo/backend/infrastructure/db"
	"github.com/graphql-todo/backend/infrastructure/todos"

	"fmt"

	"github.com/99designs/gqlgen/handler"
	"github.com/pressly/chi"
	"github.com/rs/cors"
)

const defaultPort = "8080"

func main() {
	r := chi.NewRouter()
	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	})
	r.Use(cors.Handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	var newDB db.Database
	newDB = db.NewPostgres()
	db := newDB.Open()
	defer db.Close()

	resolver := &graphql.Resolver{
		TodoUsecase: todos.NewTodoUsecase(db),
	}

	r.HandleFunc("/", handler.Playground("GraphQL playground", "/query"))
	r.HandleFunc("/query", handler.GraphQL(graphql.NewExecutableSchema(graphql.Config{Resolvers: resolver})))

	fmt.Println("アイウエオ")
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

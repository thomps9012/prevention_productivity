package main

import (
	"log"
	"net/http"
	"os"

	"thomps9012/prevention_productivity/graph"
	"thomps9012/prevention_productivity/graph/generated"
	auth "thomps9012/prevention_productivity/internal/auth"
	database "thomps9012/prevention_productivity/internal/db"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/gorilla/handlers"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()
	// router.Use(auth.Middleware())
	database.InitDB()
	defer database.CloseDB()

	router.Route("/graphql", func(router chi.Router) {
		router.Use(auth.Middleware())

		schema := generated.NewExecutableSchema(generated.Config{
			Resolvers:  &graph.Resolver{},
			Directives: generated.DirectiveRoot{},
			Complexity: generated.ComplexityRoot{},
		})

		srv := handler.NewDefaultServer(schema)
		srv.Use(extension.FixedComplexityLimit(200))
		router.Handle("/", srv)
	})

	gqlPlayground := playground.Handler("GraphQL playground", "/graphql")
	router.Get("/", gqlPlayground)
	log.Printf("Listening on localhost:%s\n", port)
	log.Printf("Visit `http://localhost:%s/graphql` in your browser", port)
	// switch below on production
	originsOk := handlers.AllowedOrigins([]string{"http://localhost:8080", "https://prevention-productivity.vercell.app"})
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
	http.ListenAndServe(":"+port, handlers.CORS(originsOk, headersOk, methodsOk)(router))
}

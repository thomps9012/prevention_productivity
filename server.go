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

	router.Route("/graphql", func(router chi.Router){
		router.Use(auth.Middleware())

		schema := generated.NewExecutableSchema(generated.Config{
			Resolvers: &graph.Resolver{},
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
log.Printf("Visit `http://localhost:%s/graphql` in your browswer", port)
panic(http.ListenAndServe(":"+port, router))
}

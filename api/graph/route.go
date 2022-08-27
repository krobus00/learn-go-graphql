package graph

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"github.com/krobus00/learn-go-graphql/api/graph/generated"
	"github.com/krobus00/learn-go-graphql/api/service"
	"go.uber.org/fx"
)

type Route struct {
	Router  *chi.Mux
	Service service.Service
}

func New(router *chi.Mux, service service.Service) *Route {
	return &Route{
		Router:  router,
		Service: service,
	}
}

func (r *Route) InitRoute() {
	r.Router.Post("/query", graphQLHandler(r.Service))
	r.Router.Get("/", playgroundQLHandler("/query"))
}

func graphQLHandler(service service.Service) http.HandlerFunc {
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &Resolver{
		Service: service,
	}}))

	return h.ServeHTTP
}

func playgroundQLHandler(endpoint string) http.HandlerFunc {
	playgroundHandler := playground.Handler("GraphQL playground", endpoint)

	return playgroundHandler
}

var Module = fx.Options(
	fx.Provide(New),
)

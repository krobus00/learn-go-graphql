//go:generate go get github.com/99designs/gqlgen@v0.17.16
//go:generate go run github.com/99designs/gqlgen generate

package graph

import (
	"github.com/krobus00/learn-go-graphql/api/service"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Service service.Service
}

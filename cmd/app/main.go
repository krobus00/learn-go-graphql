package main

import (
	"github.com/krobus00/learn-go-graphql/api/bootstrap"
	"go.uber.org/fx"
)

func main() {
	fx.New(bootstrap.AppModule).Run()
}

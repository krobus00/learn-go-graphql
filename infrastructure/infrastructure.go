package infrastructure

import (
	"github.com/go-chi/chi/v5"
	"github.com/newrelic/go-agent/v3/newrelic"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Infrastructure struct {
	fx.In

	Config      *Config
	Logger      *zap.Logger
	Database    *Database
	NewRelicAPM *newrelic.Application
	Mux         *chi.Mux
}

func NewInfrastructure() *Infrastructure {
	return &Infrastructure{}
}

var Module = fx.Options(
	fx.Provide(NewConfig),
	fx.Provide(NewLogger),
	fx.Provide(NewRelicAPM),
	fx.Provide(NewDatabase),
	fx.Provide(NewRouter),
	fx.Populate(NewInfrastructure()),
)

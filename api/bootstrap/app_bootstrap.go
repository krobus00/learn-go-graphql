package bootstrap

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/common-nighthawk/go-figure"
	"github.com/krobus00/learn-go-graphql/api/graph"
	"github.com/krobus00/learn-go-graphql/api/repository"
	"github.com/krobus00/learn-go-graphql/api/service"
	"github.com/krobus00/learn-go-graphql/infrastructure"

	"go.uber.org/fx"
)

var AppModule = fx.Options(
	infrastructure.Module,
	repository.Module,
	service.Module,
	graph.Module,
	fx.Invoke(appBootstrap),
)

func appBootstrap(
	lifecycle fx.Lifecycle,
	infra infrastructure.Infrastructure,
	handler *graph.Route,
) {

	PORT := infra.Config.AppPort

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", PORT),
		Handler: handler.Router,
	}

	appStop := func(ctx context.Context) error {
		infra.Logger.Info("Stopping Application")
		server.Shutdown(ctx)
		conn := infra.Database.DB
		conn.Close()
		infra.NewRelicAPM.Shutdown(5 * time.Second)
		return nil
	}

	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			infra.Logger.Info("Starting Application")
			figure.NewColorFigure(infra.Config.AppName, "", "purple", true).Print()
			go func() {

				err := infra.Database.DB.Ping()
				if err != nil {
					infra.Logger.Panic(err.Error())
					panic(err)
				} else {
					infra.Logger.Info("Database connected")
				}

				handler.InitRoute()

				infra.Logger.Info(fmt.Sprintf("Application running on http://0.0.0.0:%s", PORT))

				err = server.ListenAndServe()
				if err != nil {
					infra.Logger.Panic(err.Error())
				}
			}()
			return nil
		},
		OnStop: appStop,
	})
}

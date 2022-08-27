package infrastructure

import (
	"github.com/newrelic/go-agent/v3/integrations/nrzap"
	"github.com/newrelic/go-agent/v3/newrelic"
	"go.uber.org/zap"
)

func NewRelicAPM(config *Config, logger *zap.Logger) (*newrelic.Application, error) {
	if !config.NewRelicEnable {
		logger.Warn("New Relic is not enabled")
		return nil, nil
	}
	app, err := newrelic.NewApplication(
		func(c *newrelic.Config) {
			c.AppName = config.AppName
			c.License = config.NewRelicLicense
			c.Enabled = config.NewRelicEnable
			c.DistributedTracer.Enabled = true
		},
		nrzap.ConfigLogger(logger),
	)

	return app, err
}

package infrastructure

import "github.com/spf13/viper"

type Config struct {
	AppName    string `mapstructure:"APP_NAME"`
	AppPort    string `mapstructure:"APP_PORT" default:"7171"`
	AppEnv     string `mapstructure:"APP_ENV" default:"production"`
	AppTimeOut int    `mapstructure:"APP_TIMEOUT_IN_SECOND" default:"20"`

	DBDialect               string `mapstructure:"DATABASE_DIALECT"`
	DBUsername              string `mapstructure:"DATABASE_USERNAME"`
	DBPassword              string `mapstructure:"DATABASE_PASSWORD"`
	DBHost                  string `mapstructure:"DATABASE_HOST"`
	DBPort                  string `mapstructure:"DATABASE_PORT" default:"3306"`
	DBName                  string `mapstructure:"DATABASE_NAME"`
	DBMaxOpenConnection     int    `mapstructure:"DATABASE_MAX_OPEN_CONNECTION"`
	DBMaxIdleConnection     int    `mapstructure:"DATABASE_MAX_IDLE_CONNECTION_IN_SECOND"`
	DBMaxLifeTimeConnection int    `mapstructure:"DATABASE_MAX_LIFETIME_CONNECTION_IN_SECOND"`

	NewRelicEnable  bool   `mapstructure:"NEW_RELIC_ENABLE"`
	NewRelicLicense string `mapstructure:"NEW_RELIC_LICENSE"`
}

func NewConfig() (*Config, error) {
	config := &Config{}
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&config)
	return config, err
}

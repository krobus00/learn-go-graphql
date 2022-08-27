package infrastructure

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/newrelic/go-agent/v3/integrations/nrmysql"
)

type Querier interface {
	NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	Rebind(query string) string
	BindNamed(query string, arg interface{}) (string, []interface{}, error)
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
}

type Database struct {
	SqlxDB  *sqlx.DB
	Querier Querier
	DB      *sql.DB
}

func NewDatabase(config *Config) (*Database, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DBUsername,
		config.DBPassword,
		config.DBHost,
		config.DBPort,
		config.DBName,
	)
	conn, err := sqlx.Open("nrmysql", dsn)
	conn.DB.SetConnMaxIdleTime(time.Duration(config.DBMaxIdleConnection) * time.Second)
	if config.DBMaxOpenConnection > 0 {
		conn.DB.SetMaxOpenConns(config.DBMaxOpenConnection)
	}
	conn.DB.SetConnMaxLifetime(time.Duration(config.DBMaxLifeTimeConnection) * time.Second)
	return &Database{
		SqlxDB:  conn,
		Querier: conn,
		DB:      conn.DB,
	}, err
}

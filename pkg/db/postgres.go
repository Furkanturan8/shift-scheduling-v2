package db

import (
	"fmt"
	"shift-scheduling-V2/config"
	"shift-scheduling-V2/pkg/db/gorm"
	"shift-scheduling-V2/pkg/logger"
)

const (
	maxOpenConns    = 60
	connMaxLifetime = 120
	maxIdleConns    = 30
	connMaxIdleTime = 20
)

// Return new Postgresql db instance
func NewPsqlDB(c *config.Configuration, logger logger.Logger) (*gorm.Gorm, error) {
	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		c.Postgres.PostgresqlHost,
		c.Postgres.PostgresqlPort,
		c.Postgres.PostgresqlUser,
		c.Postgres.PostgresqlDbname,
		c.Postgres.PostgresqlPassword,
	)

	db, err := gorm.New(gorm.Config{
		DSN:             dataSourceName,
		MaxIdleConns:    maxIdleConns,
		MaxOpenConns:    maxOpenConns,
		MaxLifetime:     connMaxLifetime,
		ConnMaxIdleTime: connMaxIdleTime,
	}, logger)
	if err != nil {
		panic(err)
	}

	return &db, nil
}

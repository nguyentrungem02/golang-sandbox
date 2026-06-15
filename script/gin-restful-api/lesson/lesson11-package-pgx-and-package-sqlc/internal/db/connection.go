package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"trungem.com/hoc-gin/internal/config"
	"trungem.com/hoc-gin/internal/db/sqlc"
)

var DB *sqlc.Queries

func InitDB() error {
	connStr := config.NewConfig().DNS()

	configParse, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return fmt.Errorf("error parsing connection string: %v", err)
	}

	configParse.MaxConns = 50
	configParse.MinConns = 5
	configParse.MaxConnLifetime = 30 * time.Minute
	configParse.MaxConnIdleTime = 5 * time.Minute
	configParse.HealthCheckPeriod = 1 * time.Minute

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	DBPool, err := pgxpool.NewWithConfig(ctx, configParse)
	if err != nil {
		return fmt.Errorf("error creating DB pool: %v", err)
	}

	DB = sqlc.New(DBPool)

	if err := DBPool.Ping(ctx); err != nil {
		return fmt.Errorf("error pinging DB pool: %v", err)
	}

	log.Println("Connected to DB")

	return nil
}

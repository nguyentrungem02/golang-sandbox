package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"trungem.com/hoc-gin/internal/config"
)

var DB *gorm.DB

func InitDB() error {
	connStr := config.NewConfig().DNS()

	configLog := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}

	var err error
	DB, err = gorm.Open(postgres.New(postgres.Config{
		DSN: connStr,
	}), configLog)

	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	// Get generic database object sql.DB to use its functions
	sqlDB, err := DB.DB()
	if err != nil {
		return fmt.Errorf("error getting sql.DB: %w", err)
	}

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(50)

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool
	sqlDB.SetMaxIdleConns(10)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(30 * time.Minute)

	sqlDB.SetConnMaxIdleTime(5 * time.Minute)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := sqlDB.PingContext(ctx); err != nil {
		return fmt.Errorf("error pinging database: %w", err)
	}

	log.Println("Connected to database")

	return nil
}

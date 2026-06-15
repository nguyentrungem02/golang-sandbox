package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
	"trungem.com/hoc-gin/internal/config"
)

var DB *sql.DB

func InitDB() error {
	connStr := config.NewConfig().DNS()

	var err error

	// Opening a driver typically will not attempt to connect to the database.
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		// This will not be a connection error, but a DSN parse error or
		// another initialization error.
		log.Fatal("unable to use data source name", err)
	}

	DB.SetMaxIdleConns(3)                   // Số conn nhàn rõi: khi truy vấn song conn cần đóng nhưng có 1 số trường hợp nó ko đóng, nó nằm mãi trên server, ta mong muốn tối đa có 3 conn nhàn rõi đó
	DB.SetMaxOpenConns(3)                   // số conn tối đa
	DB.SetConnMaxLifetime(30 * time.Minute) // Close conn after 30 minute
	DB.SetConnMaxIdleTime(5 * time.Minute)  // Đóng conn nhàn rỗi sau 5 phút

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := DB.PingContext(ctx); err != nil {
		DB.Close()
		return fmt.Errorf("DB ping error: %w", err)
	}

	return nil
}

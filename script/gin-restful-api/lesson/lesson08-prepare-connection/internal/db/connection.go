package db

import "trungem.com/hoc-gin/internal/config"

var DB string

func InitDB() error {
	connStr := config.NewConfig().DNS()

	DB = connStr

	return nil
}

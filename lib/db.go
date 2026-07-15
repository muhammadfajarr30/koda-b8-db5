package lib

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func Conn() *pgx.Conn {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Cannot read .env file")
	}

	connStr := os.Getenv("DATABASE_URL")

	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		fmt.Println("Cannot connect to database:", err)
	}

	return conn
}

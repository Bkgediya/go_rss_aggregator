package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func Connect() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	db_url := os.Getenv("DATABASE_URL")

	pool, err := pgxpool.New(ctx, db_url)
	if err != nil {
		log.Fatal("Unable to connect to database:", err)
	}

	DB = pool

	fmt.Println("database is connecting from here")
}

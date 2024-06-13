package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
	"log"
	"os"
)

var instance *pgxpool.Pool

func DBUrl() string {
	dbUser := os.Getenv("DATABASE_USER")
	dbPass := os.Getenv("DATABASE_PASSWORD")
	dbHost := os.Getenv("DATABASE_HOST")
	dbName := os.Getenv("DATABASE_NAME")
	dbPort := os.Getenv("DATABASE_PORT")

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPass, dbName)
}

func Connect() (*pgxpool.Pool, error) {
	db, err := pgxpool.New(context.Background(), DBUrl())
	if err != nil {
		log.Fatal(err)
	}

	instance = db

	return db, nil
}

func Close() {
	instance.Close()
}

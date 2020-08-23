package database

import (
	"context"
	"fmt"
	"log"
	"os"

	// "github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Database struct {
}

func (db *Database) Setup() {
	conn := db.CreateConnection()

	initialise(context.Background(), conn)

	defer conn.Close()

	fmt.Println("Successfully created tables")
}

func (db *Database) CreateConnection() *pgxpool.Pool {
	conn, err := pgxpool.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Printf("Unable to connect to database: %v", err)
		os.Exit(1)
	}

	return conn
}

package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "dbuser"
	password = "testdb"
	dbname   = "opiniatadb"
)

func createConnection() *pgxpool.Pool {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	dbPool, err := pgxpool.Connect(context.Background(), psqlInfo)
	if err != nil {
		// fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		log.Printf("Unable to connect to database: %v", err)
		os.Exit(1)
	}

	return dbPool
}

func Create(q string) string {
	var result string

	dbPool := createConnection()

	err := dbPool.QueryRow(context.Background(), q).Scan(&result)

	if err == nil {
		fmt.Println(err)
	}

	fmt.Println(result)

	defer dbPool.Close()

	return result
}

func ExecuteQuery(s string) string {

	var result string

	dbPool := createConnection()

	err := dbPool.QueryRow(context.Background(), s).Scan(&result)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	defer dbPool.Close()

	return result
}

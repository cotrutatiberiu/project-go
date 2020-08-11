package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Database struct {
}

func (Database) Setup() {
	pool := CreateConnection()

	createAccountsTable(pool)
	createCompaniesTable(pool)

	defer pool.Close()

	fmt.Printf("Successfully created users table\n")
}

func createAccountsTable(pool *pgxpool.Pool) {
	_, err := pool.Exec(context.Background(),
		``)
	if err != nil {
		fmt.Println("Unable to create users table: ", err)
		os.Exit(1)
	}
}

func createCompaniesTable(pool *pgxpool.Pool) {
	_, err := pool.Exec(context.Background(),
		``)
	if err != nil {
		fmt.Println("Unable to create companies table: ", err)
		os.Exit(1)
	}
}

func CreateConnection() *pgxpool.Pool {
	pool, err := pgxpool.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Printf("Unable to connect to database: %v", err)
		os.Exit(1)
	}

	return pool
}

func Create(pool *pgxpool.Pool, query string) string {
	var result string

	err := pool.QueryRow(context.Background(), query).Scan(&result)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)

	return result
}

func Read(pool *pgxpool.Pool, query string) {

}

func Update() {

}

func Delete() {

}

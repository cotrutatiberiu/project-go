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
	pool := createConnection()

	createAccountsTable(pool)
	createCompaniesTable(pool)

	defer pool.Close()

	fmt.Printf("Successfully created users table\n")
}

func createAccountsTable(pool *pgxpool.Pool) {
	_, err := pool.Exec(context.Background(),
		`CREATE TABLE IF NOT EXISTS
		accounts(
		account_id SERIAL PRIMARY KEY,
		first_name VARCHAR(20) NOT NULL,
		last_name VARCHAR(20) NOT NULL,
		user_name VARCHAR(20) NOT NULL UNIQUE,
		email VARCHAR(100) NOT NULL UNIQUE,
		language VARCHAR(4) NOT NULL,
		password VARCHAR(100) NOT NULL,
		active BOOLEAN DEFAULT FALSE,
		admin_level INT DEFAULT 0,
		created TIMESTAMPTZ NOT NULL,
		updated TIMESTAMPTZ DEFAULT NULL
	  );`)
	if err != nil {
		fmt.Println("Unable to create users table: ", err)
		os.Exit(1)
	}
}

func createCompaniesTable(pool *pgxpool.Pool) {
	_, err := pool.Exec(context.Background(),
		`CREATE TABLE IF NOT EXISTS
	companies(
	company_id SERIAL PRIMARY KEY,
	account_id SERIAL NOT NULL,
	domain VARCHAR(20) NOT NULL,
	subdomain VARCHAR(20) NOT NULL,
	country VARCHAR(20) NOT NULL,
	city VARCHAR(20) DEFAULT NULL,
	language VARCHAR(4) NOT NULL,

	name VARCHAR(20) NOT NULL,
	description VARCHAR(200) NOT NULL,
	
	verified BOOLEAN DEFAULT FALSE,
	one_mark INT DEFAULT 0, 
	two_mark INT DEFAULT 0,
	three_mark INT DEFAULT 0,
	four_mark INT DEFAULT 0,
	five_mark INT DEFAULT 0,
	created TIMESTAMPTZ NOT NULL,
	updated TIMESTAMPTZ DEFAULT NULL,
	
	FOREIGN KEY(account_id) REFERENCES accounts(account_id)
	)`)
	if err != nil {
		fmt.Println("Unable to create companies table: ", err)
		os.Exit(1)
	}
}

func createConnection() *pgxpool.Pool {
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

func Read(pool *pgxpool.Pool, query string){

}

func Update(){

}

func Delete(){
	
}
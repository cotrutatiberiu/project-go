package database

import (
	"context"

	"github.com/jackc/pgx/v4"
)

const schema = `
	CREATE TABLE IF NOT EXISTS accounts(
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
	);

	CREATE TABLE IF NOT EXISTS companies (
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
	);
`

// Init create database schema.
func Init(ctx context.Context, conn *pgx.Conn) error {
	_, err := conn.Exec(ctx, schema)
	return err
}

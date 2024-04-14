package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBname   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s dbuser=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.DBname, cfg.Password, cfg.SSLMode))

	if err != nil {
		fmt.Println("failed to open db")
		return nil, err
	}
	defer db.Close()

	createTableUsers := `
		DROP TABLE IF EXISTS users;
		CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		username VARCHAR(255) NOT NULL,
		first_name VARCHAR(255) NOT NULL,
		last_name VARCHAR(255) NOT NULL,
		email VARCHAR(255) NOT NULL,
		password VARCHAR(255) NOT NULL,
		created_at DATE NOT NULL,
		UNIQUE (email)		
	)`

	_, err = db.Exec(createTableUsers)
	if err != nil {
		fmt.Println("failed to create table users")
		return nil, err
	}

	createTableTransactions := `
		DROP TABLE IF EXISTS transactions;
		CREATE TABLE IF NOT EXISTS transactions (
		id SERIAL PRIMARY KEY,
		amount FLOAT NOT NULL,
		currency VARCHAR(3) NOT NULL,
		type VARCHAR(7) NOT NULL,
		category VARCHAR(255) NOT NULL,
		description VARCHAR(255) NOT NULL,
		date DATE NOT NULL,
		user_id INT NOT NULL,
		FOREIGN KEY (user_id) REFERENCES users(id)
	)`

	_, err = db.Exec(createTableTransactions)
	if err != nil {
		fmt.Println("failed to create table transactions")
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("failed to connect to db")
		return nil, err
	}

	return db, nil
}

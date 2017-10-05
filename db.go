package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// Engine description
type Engine struct {
	DB *sql.DB
}

// Initialize description
func (e *Engine) Initialize(user, password, dbname, host string) {
	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=disable", user, password, dbname, host)

	var err error
	e.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
}

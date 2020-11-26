package main

import (
	"bank/db"
	"bank/pkg/core/services"
	"database/sql"
	"log"
	_ "mattn/go-sqlite3"
)

func main() {
	database, err := sql.Open("sqlite3", "test")
	if err != nil {
		log.Fatalf("Err is %e")
	}
	db.Dbinit(database)

	Start(database)
}

func Start(database *sql.DB){
	for{
		login, password := services.Authorization()
		services.Login(database, login, password)
	}
}
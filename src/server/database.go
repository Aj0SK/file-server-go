package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func insertUser(db *sql.DB, user string, password string) bool {
	stmt, err := db.Prepare("INSERT INTO users(username,password) VALUES(?,?)")
	if err != nil {
		log.Fatal(err)
		return false
	}
	_, err = stmt.Exec(user, hashMD5(password))
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

func ConnectDB() *sql.DB {
	db, err := sql.Open("mysql", "Aj0:Poklopadresa123@/test")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Connection Established")
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("Can't ping database.")
	}

	return db
}

func DisconnectDB(db *sql.DB) {
	defer db.Close()
}

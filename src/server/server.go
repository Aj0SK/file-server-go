package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, Your files %s:\n%s", r.URL.Path[1:], listFiles("../../storage/"))
}

func main() {

	db := ConnectDB()

	insertUser(db, "user12", "ee")
	insertUser(db, "user13", "kk")

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))

	DisconnectDB(db)
}

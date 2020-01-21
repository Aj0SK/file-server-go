package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db = ConnectDB()

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, Your files %s:\n%s", r.URL.Path[1:], listFiles("../../storage/"))
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

//curl -i -X POST -H 'Content-Type: application/json' -d '{ "username": "testUser", "password": "abc" }' http://localhost:8080/signup
func Signup(w http.ResponseWriter, r *http.Request) {
	creds := &Credentials{}
	err := json.NewDecoder(r.Body).Decode(creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	succ := insertUser(db, creds.Username, creds.Password)
	if succ != true {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

//curl -i -X POST -H 'Content-Type: application/json' -d '{ "username": "testUser", "password": "abc" }' http://localhost:8080/signin
func Signin(w http.ResponseWriter, r *http.Request) {
	creds := &Credentials{}
	err := json.NewDecoder(r.Body).Decode(creds)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Get the existing entry present in the database for the given username
	res := authUser(db, creds.Username, hashMD5(creds.Password))
	if res != true {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	return
}

func main() {

	http.HandleFunc("/", handler)
	http.HandleFunc("/signup", Signup)
	http.HandleFunc("/signin", Signin)
	log.Fatal(http.ListenAndServe(":8080", nil))

	DisconnectDB(db)
}

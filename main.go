package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	//	"github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	// "strconv"
)

const (
	dbDriver = "mysql"
	dbUser   = "dbUser"
	dbPass   = "dbPass"
	dbName   = "gocrud_app"
)

func main() {
	// create a new router
	r := mux.NewRouter()

	// define http routes using the router
	r.HandleFunc("/user", createUserHandler).Methods("POST")
	//	r.HandleFunc("/user/{id}", getUserHandler).Methods("GET")
	//	r.HandleFunc("/user{id}", updateUserHandler).Methods("PUT")
	//	r.HandleFunc("/user{id}", deleteUserHandler).Methods("DELETE")

	// start http port on the server 8090
	log.Println("Server listening on the port 8090")
	log.Fatal(http.ListenAndServe(":8090", r))
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open(dbDriver, dbUser+":"+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Parse JSON data from the request body
	var user User
	json.NewDecoder(r.Body).Decode(&user)

	CreateUser(db, user.Name, user.Email)
	if err != nil {
		http.Error(w, "Failed to create  user", http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "User created  successfully")
}

func CreateUser(db *sql.DB, name, email string) error {
	query := "INSERT INTO  users (name, email) VALUES  (?, ?)"
	_, err := db.Exec(query, name, email)
	if err != nil {
		return err
	}
	return nil
}

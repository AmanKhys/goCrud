package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	//	"github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
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

// ////////////////////////////////////////////////////////////////
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

// //////////////////////////////////////////////////////////////////////
type User struct {
	ID    int
	Name  string
	Email string
}

// /////////////////////////////////////////////////////////////////////////
func getUserHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Get the 'id' parameter from the URL
	vars := mux.Vars(r)
	idStr := vars["id"]

	// Convert 'id' to an integer
	userID, err := strconv.Atoi(idStr)

	// Call the GetUser function ot fetch the user data from the database
	user, err := GetUser(db, userID)
	if err != nil {
		http.Error(w, "User  not found", http.StatusNotFound)
		return
	}

	// Convert the user object to JSON and send it in the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func GetUser(db *sql.DB, id int) (*User, error) {
	query := "SELECT * FROM users WHERE id = ?"
	row := db.QueryRow(query, id)

	user := &User{}
	err := row.Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

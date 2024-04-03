package GoCrud

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/go-sql-driver/mysql"
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
	r.HandleFunc("/user/{id}", getUserHandler).Methods("GET")
	r.HandleFunc("/user{id}", updateUserHandler).Methods("PUT")
	r.HandleFunc("/user{id}", deleteUserHandler).Methods("DELETE")

	// start http port on the server 8090
	log.Println("Server listening on the port 8090")
	log.Fatal(http.ListenAndServe(":8090", r))
}

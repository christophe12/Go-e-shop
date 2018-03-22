package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// App defines our app
type app struct {
	Router *mux.Router
	DB     *gorm.DB
}

//Initialize creates db connection and initialize mux router
func (a *app) Initialize(userr string, password string, dbname string) {
	connectionString := fmt.Sprintf("%s:%s@/%s", userr, password, dbname)
	var err error
	a.DB, err = gorm.Open("mysql", connectionString)

	// migrate db
	migrateSchema(a)

	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()
	a.initialiazeRoutes()

}

//Run starts the application
func (a *app) Run(addr string) {
	// this will simply start thapplication
	log.Printf("App running on:http://localhost:8080")
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func respondWithError(w http.ResponseWriter, code int, err []string) {
	response := ResponseBuilder{
		Status:     "failure",
		StatusCode: code,
		Errors:     err,
	}
	responseWrite(w, response)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response := ResponseBuilder{
		Status:     "success",
		StatusCode: code,
		Data:       payload,
	}
	responseWrite(w, response)
}

func responseWrite(w http.ResponseWriter, response ResponseBuilder) {
	jsonResponse, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)
	w.Write(jsonResponse)
}

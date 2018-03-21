package main

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (a *app) getProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var p Product
	if err := a.DB.Where("id = ?", id).First(&p).Error; err != nil {
		respondWithError(w, http.StatusNotFound, "Product not found")
		return
	}

	respondWithJSON(w, http.StatusOK, p)
}

// func (a *app) createProduct(w http.ResponseWriter, r *http.Request) {
// 	var p Product
// 	decoder := json.NewDecoder(r.Body)
// 	err := decoder.Decode(&p)
// 	if err != nil {
// 		respondWithError(w, http.StatusBadRequest, err.Error())
// 		return
// 	}
// 	defer r.Body.Close()

// 	// validations
// 	if intIsEmpty(p.UserID) {
// 		respondWithError(w, http.StatusBadRequest, "User id is required")
// 		return
// 	}

// 	if intIsEmpty(p.StoreID) {
// 		respondWithError(w, http.StatusBadRequest, "Store id is required")
// 		return
// 	}

// 	if stringIsEmpty(p.Name) {
// 		respondWithError(w, http.StatusBadRequest, "Store name is required")
// 		return
// 	}

// 	// hashing password beforesave
// 	u.hashUserPassword()

// 	a.DB.NewRecord(p)

// 	a.DB.Create(&p)

// 	// user was not saved
// 	if a.DB.NewRecord(p) != false {
// 		respondWithError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	respondWithJSON(w, http.StatusCreated, p)

// }

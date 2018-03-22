package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (a *app) getCurrency(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var c Currency
	if err := a.DB.Where("id = ?", id).First(&c).Error; err != nil {
		respondWithError(w, http.StatusNotFound, []string{"Product not found"})
		return
	}

	respondWithJSON(w, http.StatusOK, c)
}

func (a *app) createCurrency(w http.ResponseWriter, r *http.Request) {
	var c Currency
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&c)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, []string{err.Error()})
		return
	}
	defer r.Body.Close()

	// validations
	if stringIsEmpty(c.Name) || stringIsEmpty(c.Code) {
		respondWithError(w, http.StatusBadRequest, []string{"Currency Name and Code are required"})
		return
	}

	a.DB.NewRecord(c)

	a.DB.Create(&c)

	// currency was not saved
	if a.DB.NewRecord(c) != false {
		respondWithError(w, http.StatusInternalServerError, []string{err.Error()})
		return
	}

	respondWithJSON(w, http.StatusCreated, c)

}

package main

import (
	"encoding/json"
	"net/http"
)

func (a *app) createRole(w http.ResponseWriter, r *http.Request) {
	var role Role
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&role)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, []string{err.Error()})
		return
	}
	defer r.Body.Close()

	// validations
	var validationErrors []string
	if stringIsEmpty(role.Name) {
		validationErrors = append(validationErrors, "Role name is required")
	}

	a.DB.NewRecord(role)

	a.DB.Create(&role)

	if a.DB.NewRecord(role) != false {
		validationErrors = append(validationErrors, "We could not save your role. Please try again!")
	}

	respondWithJSON(w, http.StatusCreated, role)
}

func (a *app) getAllRoles(w http.ResponseWriter, r *http.Request) {
	roles := []Role{}
	a.DB.Find(&roles)
	respondWithJSON(w, http.StatusOK, roles)
}

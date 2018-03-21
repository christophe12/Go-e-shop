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
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	// validations
	if stringIsEmpty(role.Name) {
		respondWithError(w, http.StatusBadRequest, "Role name is required")
		return
	}

	a.DB.NewRecord(role)

	a.DB.Create(&role)

	if a.DB.NewRecord(role) != false {
		respondWithError(w, http.StatusInternalServerError, "We could not save your role. Please try again!")
		return
	}

	respondWithJSON(w, http.StatusCreated, role)
}

func (a *app) getAllRoles(w http.ResponseWriter, r *http.Request) {
	roles := []Role{}
	a.DB.Find(&roles)
	respondWithJSON(w, http.StatusOK, roles)
}

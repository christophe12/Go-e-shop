package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (a *app) getUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var u User
	a.DB.Where("id = ?", id).First(&u)

	if u.ID == 0 {
		respondWithError(w, http.StatusNotFound, "User not found")
		return
	}

	respondWithJSON(w, http.StatusOK, u.embedUserItem(a))
}

func (a *app) createUser(w http.ResponseWriter, r *http.Request) {
	var u User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&u)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	// validations
	if stringIsEmpty(u.Name) || stringIsEmpty(u.Email) {
		respondWithError(w, http.StatusBadRequest, "Name and Email are required")
		return
	}

	if stringIsEmpty(u.Password) {
		respondWithError(w, http.StatusBadRequest, "The password is required")
		return
	}

	if intIsEmpty(u.RoleID) {
		respondWithError(w, http.StatusBadRequest, "Role id is required")
		return
	}

	if !isValidEmail(u.Email) {
		respondWithError(w, http.StatusBadRequest, "Invalid Email: "+u.Email)
		return
	}

	// hashing password beforesave
	u.hashUserPassword()

	a.DB.NewRecord(u)

	a.DB.Create(&u)

	// user was not saved
	if a.DB.NewRecord(u) != false {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, u)

}

func (a *app) getAllUsers(w http.ResponseWriter, r *http.Request) {
	users := []User{}
	a.DB.Find(&users)
	respondWithJSON(w, http.StatusOK, embedUserCollection(users, a))
}

func (a *app) updateUser(w http.ResponseWriter, r *http.Request) {
	var newUser User
	var oldUser User

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newUser)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	// first check if the user exists
	if err := a.DB.Where("id = ?", newUser.ID).First(&oldUser).Error; err != nil {
		respondWithError(w, http.StatusNotFound, "We can't find the user you are trying to update!")
		return
	}

	if !stringIsEmpty(newUser.Name) {
		oldUser.Name = newUser.Name
	}

	if !stringIsEmpty(newUser.Email) {
		oldUser.Email = newUser.Email
	}

	if !intIsEmpty(newUser.RoleID) {
		oldUser.RoleID = newUser.RoleID
	}

	if err := a.DB.Save(&oldUser).Error; err != nil {
		respondWithError(w, http.StatusInternalServerError, "We couldn't update your store")
		return
	}

	respondWithJSON(w, http.StatusOK, oldUser)

}

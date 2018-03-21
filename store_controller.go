package main

import (
	"encoding/json"
	"net/http"
)

func (a *app) createStore(w http.ResponseWriter, r *http.Request) {
	var store Store
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&store)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	// validations
	if intIsEmpty(store.UserID) {
		respondWithError(w, http.StatusBadRequest, "The Store id for this store is required")
		return
	}

	if stringIsEmpty(store.Name) {
		respondWithError(w, http.StatusBadRequest, "The store name is required")
		return
	}

	a.DB.NewRecord(store)

	a.DB.Create(&store)

	// Store was not saved
	if a.DB.NewRecord(store) != false {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, store)

}

func (a *app) getAllStores(w http.ResponseWriter, r *http.Request) {
	stores := []Store{}
	a.DB.Find(&stores)
	respondWithJSON(w, http.StatusOK, stores)
}

func (a *app) updateStore(w http.ResponseWriter, r *http.Request) {
	var newStore Store
	var oldStore Store

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newStore)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	// first check if the Store exists
	if err := a.DB.Where("id = ?", newStore.ID).First(&oldStore).Error; err != nil {
		respondWithError(w, http.StatusNotFound, "We can't find the store you are trying to update!")
		return
	}

	if !stringIsEmpty(newStore.Name) {
		oldStore.Name = newStore.Name
	}

	if !stringIsEmpty(newStore.Description) {
		oldStore.Description = newStore.Description
	}

	if err := a.DB.Save(&oldStore).Error; err != nil {
		respondWithError(w, http.StatusInternalServerError, "We couldn't update your store")
		return
	}

	respondWithJSON(w, http.StatusOK, oldStore)
}

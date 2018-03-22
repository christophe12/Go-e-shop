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
		respondWithError(w, http.StatusBadRequest, []string{err.Error()})
		return
	}
	defer r.Body.Close()

	// validations
	var validationErrors []string
	if intIsEmpty(store.UserID) {
		validationErrors = append(validationErrors, "The Store id for this store is required")
	}

	if stringIsEmpty(store.Name) {
		validationErrors = append(validationErrors, "The store name is required")
	}

	a.DB.NewRecord(store)

	a.DB.Create(&store)

	// Store was not saved
	if a.DB.NewRecord(store) != false {
		respondWithError(w, http.StatusInternalServerError, []string{err.Error()})
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
		respondWithError(w, http.StatusBadRequest, []string{err.Error()})
		return
	}
	defer r.Body.Close()

	// first check if the Store exists
	if err := a.DB.Where("id = ?", newStore.ID).First(&oldStore).Error; err != nil {
		respondWithError(w, http.StatusNotFound, []string{"We can't find the store you are trying to update!"})
		return
	}

	if !stringIsEmpty(newStore.Name) {
		oldStore.Name = newStore.Name
	}

	if !stringIsEmpty(newStore.Description) {
		oldStore.Description = newStore.Description
	}

	if err := a.DB.Save(&oldStore).Error; err != nil {
		respondWithError(w, http.StatusInternalServerError, []string{"We couldn't update your store"})
		return
	}

	respondWithJSON(w, http.StatusOK, oldStore)
}

package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (a *app) getProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var p Product
	if err := a.DB.Where("id = ?", id).First(&p).Error; err != nil {
		respondWithError(w, http.StatusNotFound, []string{"Product not found"})
		return
	}
	// manipulations
	p.retrievePrice()

	respondWithJSON(w, http.StatusOK, p.embedProductItem(a))
}

func (a *app) createProduct(w http.ResponseWriter, r *http.Request) {
	var p Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&p)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, []string{err.Error()})
		return
	}
	defer r.Body.Close()

	// validations
	var validationErrors []string

	if intIsEmpty(p.UserID) {
		validationErrors = append(validationErrors, "User id is required")
	}

	if intIsEmpty(p.StoreID) {
		validationErrors = append(validationErrors, "Store id is required")
	}

	if stringIsEmpty(p.Name) {
		validationErrors = append(validationErrors, "Product name is required")
	}

	if intIsEmpty(p.CurrencyID) {
		validationErrors = append(validationErrors, "Product currency id is required")
	}

	if p.Price == 0 {
		validationErrors = append(validationErrors, "Product price is required")
	}

	if validationErrors != nil {
		respondWithError(w, http.StatusBadRequest, validationErrors)
		return
	}

	// manipulations
	p.createSlug()
	p.storePrice()

	a.DB.NewRecord(p)

	a.DB.Create(&p)

	// product was not saved
	if a.DB.NewRecord(p) != false {
		respondWithError(w, http.StatusInternalServerError, []string{err.Error()})
		return
	}

	respondWithJSON(w, http.StatusCreated, p.embedProductItem(a))

}

func (a *app) getAllProducts(w http.ResponseWriter, r *http.Request) {
	products := []Product{}
	a.DB.Find(&products)

	respondWithJSON(w, http.StatusOK, embedProductCollection(retrievePrices(products), a))
}

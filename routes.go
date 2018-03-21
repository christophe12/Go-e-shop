package main

func (a *app) initialiazeRoutes() {

	// users routes
	a.Router.HandleFunc("/users", a.getAllUsers).Methods("GET")
	a.Router.HandleFunc("/users/{id}", a.getUser).Methods("GET")
	a.Router.HandleFunc("/users/create", a.createUser).Methods("POST")
	a.Router.HandleFunc("/users/update", a.updateUser).Methods("POST")

	// roles routes
	a.Router.HandleFunc("/roles", a.getAllRoles).Methods("GET")
	a.Router.HandleFunc("/roles/create", a.createRole).Methods("POST")

	// stores routes
	a.Router.HandleFunc("/stores", a.getAllStores).Methods("GET")
	a.Router.HandleFunc("/stores/create", a.createStore).Methods("POST")

	// products routes
	a.Router.HandleFunc("/products/{id}", a.getProduct).Methods("GET")
}

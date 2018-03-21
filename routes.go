package main

func (a *app) initialiazeRoutes() {
	// a.Router.HandleFunc("/", a.index).Methods("GET")

	// users routes
	a.Router.HandleFunc("/users", a.getAllUsers).Methods("GET")
	a.Router.HandleFunc("/users/{id}", a.getUser).Methods("GET")
	a.Router.HandleFunc("/users/create", a.createUser).Methods("POST")

	// roles routes
	a.Router.HandleFunc("/roles", a.getAllRoles).Methods("GET")
	a.Router.HandleFunc("/roles/create", a.createRole).Methods("POST")
}

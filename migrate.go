package main

func migrateSchema(a *app) {

	/**
	*
	* create tables
	*
	 */

	// roles table
	if !a.DB.HasTable("roles") {
		var r Role
		a.DB.CreateTable(&r)
	}

	// users table
	if !a.DB.HasTable("users") {
		var u User
		a.DB.CreateTable(&u)
	}

	// products table
	if !a.DB.HasTable("products") {
		var p Product
		a.DB.CreateTable(&p)
	}

	// stores table
	if !a.DB.HasTable("stores") {
		var s Store
		a.DB.CreateTable(&s)
	}

	// currencies table
	if !a.DB.HasTable("currencies") {
		var c Currency
		a.DB.CreateTable(&c)
	}

}

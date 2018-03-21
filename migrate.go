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

}

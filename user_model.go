package main

//User model
type User struct {
	ID       int    `gorm:"AUTO_INCREMENT;not null;PRIMARY_KEY" json:"id"`
	RoleID   int    `gorm:"type:int(11);not null" json:"role_id"`
	Name     string `gorm:"type:varchar(255);not null" json:"name"`
	Email    string `gorm:"type:varchar(255);unique;not null" json:"email"`
	Password string `gorm:"type:varchar(255);not null" json:"password"`
}

//UserEmbeds defines data to be associated with a user
type UserEmbeds struct {
	U User    `json:"user"`
	S []Store `json:"stores"`
}

// mutators
func (u *User) hashUserPassword() {
	(*u).Password = HashAndSalt((*u).Password)
}

// include necessary data to return with the User
func (u User) embedUserItem(a *app) UserEmbeds {
	stores := []Store{}
	a.DB.Where("user_id = ?", u.ID).Find(&stores)
	return UserEmbeds{
		U: u,
		S: stores,
	}
}

func embedUserCollection(users []User, a *app) []UserEmbeds {
	embeds := []UserEmbeds{}
	for _, user := range users {
		embeds = append(embeds, user.embedUserItem(a))
	}
	return embeds
}

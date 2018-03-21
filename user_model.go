package main

//User model
type User struct {
	ID       int    `gorm:"AUTO_INCREMENT;not null;PRIMARY_KEY" json:"id"`
	RoleID   int    `gorm:"type:int(11);not null" json:"role_id"`
	Name     string `gorm:"type:varchar(255);not null" json:"name"`
	Email    string `gorm:"type:varchar(255);unique;not null" json:"email"`
	Password string `gorm:"type:varchar(255);not null" json:"password"`
}

// mutators

package main

//Role model
type Role struct {
	ID   int    `gorm:"AUTO_INCREMENT;not null;PRIMARY_KEY" json:"id"`
	Name string `gorm:"type:varchar(255);not null" json:"name"`
}

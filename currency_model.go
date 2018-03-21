package main

//Currency model
type Currency struct {
	ID   int    `gorm:"AUTO_INCREMENT;not null;PRIMARY_KEY" json:"id"`
	Name string `gorm:"type:varchar(255);not null" json:"name"`
	Code string `gorm:"type:varchar(25);not null" json:"code"`
}

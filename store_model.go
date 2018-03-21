package main

//Store model
type Store struct {
	ID          int    `gorm:"AUTO_INCREMENT;not null;PRIMARY_KEY" json:"id"`
	UserID      int    `gorm:"type:int(11);not null" json:"user_id"`
	Name        string `gorm:"type:varchar(255);not null;unique" json:"name"`
	Description string `gorm:"type:text" json:"description"`
}

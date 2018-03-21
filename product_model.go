package main

//Product model
type Product struct {
	ID          int    `gorm:"AUTO_INCREMENT;not null;PRIMARY_KEY" json:"id"`
	UserID      int    `gorm:"type:int(11);not null" json:"user_id"`
	StoreID     int    `gorm:"type:int(11);not null" json:"store_id"`
	Name        string `gorm:"type:varchar(255);not null" json:"name"`
	Description string `gorm:"type:text" json:"description"`
	Slug        string `gorm:"type:varchar(255);unique;not null" json:"slug"`
	Price       int64  `gorm:"type:int(64);not null" json:"price"`
	SalePrice   int64  `gorm:"type:int(64)" json:"sale_price"`
	CurrencyID  int    `gorm:"type:int(11);not null" json:"currency_id"`
	OnSale      int    `gorm:"type:int(11);default:0" json:"on_sale"`
	Status      string `gorm:"type:varchar(255);not null;default:'draft'"`
}

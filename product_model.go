package main

import (
	"strings"
)

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

//ProductEmbeds defines data to be associated with a product
type ProductEmbeds struct {
	P Product `json:"product"`
	S Store   `json:"store"`
	U User    `json:"user"`
}

func (p Product) embedProductItem(a *app) ProductEmbeds {
	var store Store
	a.DB.Where("id = ?", p.StoreID).First(&store)
	var user User
	a.DB.Where("id = ?", p.UserID).First(&user)

	return ProductEmbeds{
		P: p,
		S: store,
		U: user,
	}
}

func embedProductCollection(products []Product, a *app) []ProductEmbeds {
	embeds := []ProductEmbeds{}
	for _, product := range products {
		embeds = append(embeds, product.embedProductItem(a))
	}
	return embeds
}

//mutators

func (p *Product) createSlug() {
	(*p).Slug = strings.Join(strings.Split(strings.ToLower((*p).Name), " "), "-")
}

func (p *Product) storePrice() {
	(*p).Price = (*p).Price / 100

	if (*p).SalePrice != 0 {
		(*p).SalePrice = (*p).SalePrice / 100
	}
}

func (p *Product) retrievePrice() {
	(*p).Price = (*p).Price * 100

	if (*p).SalePrice != 0 {
		(*p).SalePrice = (*p).SalePrice * 100
	}
}

func retrievePrices(products []Product) []Product {
	newProducts := []Product{}
	for _, p := range products {
		p.retrievePrice()
		newProducts = append(newProducts, p)
	}
	return newProducts
}

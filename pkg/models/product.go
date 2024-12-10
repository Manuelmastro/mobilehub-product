package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	//ID uint `gorm:"primary key" json:"id"`
	//CategoryID  uint     `json:"category_id" validate:"required"`
	CategoryName string  `json:"category_name"`
	ProductName  string  `json:"product_name"`
	Description  string  `json:"product_description"`
	ImageUrl     string  `json:"product_imageUrl"`
	Price        float64 `json:"price"`
	Stock        int32   `json:"stock"`
	//Popular              bool     `gorm:"type:boolean;default:false" json:"popular" validate:"required"`
	//Size                 string   `gorm:"type:varchar(10); check:size IN ('Medium', 'Small', 'Large')" json:"size" validate:"required,oneof=Medium Small Large"`
	//HasOffer             bool `gorm:"default:false"`
	//OfferDiscountPercent uint `gorm:"default:0"`
}

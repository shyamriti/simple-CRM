package models

import "gorm.io/gorm"

type Person struct {
	gorm.Model
	Name    string `json:"name" gorm:"column:name"`
	Email   string `json:"email" gorm:"column:email"`
	Company string `json:"company" gorm:"column:company"`
	Phone   int    `json:"phone" gorm:"column:phone"`
}

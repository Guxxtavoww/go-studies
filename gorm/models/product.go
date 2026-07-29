package models

type Product struct {
	Id    int `gorm:"primaryKey"`
	Name  string
	Price float64
}
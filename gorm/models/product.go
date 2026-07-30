package models

type EntityId int

type Product struct {
	Id    EntityId `gorm:"primaryKey"`
	Name  string
	Price float64
}
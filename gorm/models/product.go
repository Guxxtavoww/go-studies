package models

type Product struct {
	Id    EntityId `gorm:"primaryKey"`
	Name  string
	Price float64
}

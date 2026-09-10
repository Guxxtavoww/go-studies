package main

import (
	"fmt"
	"log"

	"gorm/database"
	"gorm/models"
	"gorm/start"
	"gorm/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func getProductById(db *gorm.DB, productId models.EntityId) *models.Product {
	var product models.Product

	db.First(&product, productId)

	return &product
}

func getProductsByName(db *gorm.DB, name string) *[]models.Product {
	var products []models.Product

	// 	Why three %?
	// %% → literal %
	// %s → insert the string
	db.Where("name ILIKE ?", fmt.Sprintf("%%%s%%", name)).Find(&products)

	return &products
}

func main() {
	env := start.AppSetup()

	dsn := env.GetDbConnectionDsn()

	db := utils.UnrwrapError(gorm.Open(postgres.Open(dsn), &gorm.Config{}))

	utils.UnrwrapError(db.DB()).Ping()

	if err := database.RunMigrations(db); err != nil {
		log.Fatalf("failed to run migrations: %v", err)
	}

	// Create One
	// db.Create(&Product{
	// 	Name:  "Plug do denix",
	// 	Price: 669.69,
	// })

	// Create Many
	// productsCreateMany := []models.Product{
	// 	{Name: "Plug do Estevam", Price: 2022},
	// 	{Name: "Fodase", Price: 12121},
	// 	{Name: "MMMMMM", Price: 23030303},
	// }
	// db.Create(&productsCreateMany)

	// Get products by name
	// products := getProductsByName(db, "Plug")
	// fmt.Println(products)

	// Gets all products
	// var products []models.Product
	// db.Find(&products)
	// for index, product := range products {
	// 	fmt.Println(index, product)
	// }

	// gets all products where the price is lagrger than 2022
	// var products []models.Product
	// db.Where("price > ?", 2022).Find(&products)


	// Lock Pessimista (Pessimistic Lock) A ideia central: "assumo que vai dar conflito, então bloqueio antes de qualquer coisa dar errado."
	tx := db.Begin()

	var product models.Product

	err := tx.Debug().Clauses(clause.Locking{Strength: "UPDATE"}).First(&product, 1).Error

	if err != nil {
		tx.Rollback()

		return
	}

	product.Price = 9999

	tx.Debug().Save(&product)

	tx.Commit()
}

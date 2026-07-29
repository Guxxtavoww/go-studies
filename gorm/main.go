package main

import (
	"fmt"

	"gorm/models"
	"gorm/start"
	"gorm/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	env := start.AppSetup()

	dsn := fmt.Sprintf(
		"host=localhost user=%s password=%s dbname=%s port=%s sslmode=disable",
		env.DB_USER,
		env.DATABASE_ROOT_PASSWORD,
		env.DATABASE_DATABASE_NAME,
		env.DB_PORT,
	)

	db := utils.UnrwrapError(gorm.Open(postgres.Open(dsn), &gorm.Config{}))

	utils.UnrwrapError(db.DB()).Ping()

	db.AutoMigrate(&models.Product{})

	// db.Create(&Product{
	// 	Name:  "Plug do denix",
	// 	Price: 669.69,
	// })

	productsCreateMany := []models.Product{
		{Name: "Plug do Estevam", Price: 2022},
		{Name: "Fodase", Price: 12121},
		{Name: "MMMMMM", Price: 23030303},
	}

	db.Create(&productsCreateMany)
}

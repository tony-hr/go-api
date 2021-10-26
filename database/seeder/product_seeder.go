package seeder

import (
	"fmt"
	"go-api/models"

	"github.com/jinzhu/gorm"
)

func ProductSeeder(db *gorm.DB) {
	var products []models.Product

	var sample1 = models.Product{
		Name:   "Kaos Pria XL - Hitam",
		Type:   "TSHIRT",
		Price:  78000,
		Stock:  34,
		Images: "kaos-pria-black.png",
	}
	products = append(products, sample1)

	var sample2 = models.Product{
		Name:   "Daster Ibu Hamil Jumbo (Import)",
		Type:   "TSHIRT",
		Price:  94500,
		Stock:  27,
		Images: "daster-emak2.jpg",
	}
	products = append(products, sample2)

	for _, pro := range products {
		if err := db.Create(&pro).Error; err != nil {
			fmt.Println(err.Error())

		}
		fmt.Printf("Product has been created\n")
	}
}

package seeder

import (
	"fmt"
	"go-api/models"
	"log"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

func UserSeeder(db *gorm.DB) {
	var users []models.User

	var sampleAdmin = models.User{
		Username:    "admin",
		Fullname:    "Admin",
		Email:       "admin@gmail.com",
		Password:    hashPasswordUser("password"),
		PhoneNumber: "081122334455",
	}
	users = append(users, sampleAdmin)

	var sampleTony = models.User{
		Username:    "tony",
		Fullname:    "Tony HR",
		Email:       "tony@email.com",
		Password:    hashPasswordUser("passtony"),
		PhoneNumber: "081299333",
	}
	users = append(users, sampleTony)

	for _, user := range users {
		if err := db.Create(&user).Error; err != nil {
			fmt.Println(err.Error())

		}
		fmt.Printf("User has been created\n")
	}
}

func hashPasswordUser(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)

	if err != nil {
		log.Fatalln(err.Error())
	}
	return string(bytes)
}

package database

import (
	"go-api/config"
	seed "go-api/database/seeder"
	"go-api/models"

	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func InitDB() *gorm.DB {
	conf := config.GetConfig()

	dbHost := conf.GetString("database.host")
	dbPort := conf.GetString("database.port")
	dbUser := conf.GetString("database.username")
	dbPass := conf.GetString("database.password")
	dbName := conf.GetString("database.dbname")

	if "" == dbHost || "" == dbPort || "" == dbUser || "" == dbName {
		log.Fatalln("Database credentials not define.")
	}

	db, err := gorm.Open("mysql", dbUser+":"+dbPass+"@tcp("+dbHost+":"+dbPort+")/"+dbName+"?parseTime=True&loc=Local")
	if err != nil {
		log.Fatalln("Can't connect to MySQL DB. ", err)
		panic(err)
	}

	db.DB().SetMaxIdleConns(3)
	db.LogMode(true)

	return db
}

func MigrateDDL(db *gorm.DB) error {
	err := db.AutoMigrate(
		&models.User{},
		&models.Product{},
	)

	return err.Error
}

func Seeder(db *gorm.DB) {
	seed.UserSeeder(db)
	seed.ProductSeeder(db)
}

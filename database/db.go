package database

import (
	"fmt"
	"go-api/config"
	seed "go-api/database/seeder"
	"go-api/models"

	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func InitDB() *gorm.DB {
	conf := config.GetConfig()

	dbHost := conf.GetString("database.host")
	dbPort := conf.GetString("database.port")
	dbUser := conf.GetString("database.username")
	dbPass := conf.GetString("database.password")
	dbName := conf.GetString("database.dbname")
	dbTimezone := conf.GetString("database.timezone")

	if "" == dbHost || "" == dbPort || "" == dbUser || "" == dbName || "" == dbTimezone {
		log.Fatalln("Database credentials not define.")
	}

	dns := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s TimeZone=%s", dbHost, dbPort, dbUser, dbName, dbPass, dbTimezone)
	psqlConn, err := gorm.Open("postgres", dns)

	err = psqlConn.DB().Ping()
	if err != nil {
		log.Fatalln("Can't connect to Postgre DB. ", err)
		panic(err)
	}

	psqlConn.DB().SetMaxIdleConns(3)
	psqlConn.LogMode(true)

	return psqlConn
}

func MigrateDDL(db *gorm.DB) error {
	err := db.AutoMigrate(
		&models.User{},
	)

	return err.Error
}

func Seeder(db *gorm.DB) {
	seed.UserSeeder(db)
}

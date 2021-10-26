package main

import (
	ProductRepo "go-api/app/product/repository"
	ProductService "go-api/app/product/service"
	UserRepo "go-api/app/user/repository"
	UserService "go-api/app/user/service"
	"go-api/config"
	"go-api/database"
	"go-api/routes"
	"os"
)

func main() {
	conf := config.GetConfig()
	db := database.InitDB()
	router := routes.Init()

	dbEvent := os.Getenv("DBEVENT")
	if dbEvent == "migrate" {
		database.MigrateDDL(db)
		database.Seeder(db)
	}

	//User Endpoint
	userRepo := UserRepo.NewUserRepository(db)
	userService := UserService.NewUserService(userRepo)
	routes.NewUserHandler(router, userService)

	//Product Endpoint
	productRepo := ProductRepo.NewProductRepository(db)
	productService := ProductService.NewProductService(productRepo)
	routes.NewProductHandler(router, productService)

	router.Logger.Fatal(router.Start(":" + conf.GetString("app.port")))
}

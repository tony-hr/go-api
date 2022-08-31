package main

import (
	JobService "go-api/app/job/service"
	UserRepo "go-api/app/user/repository"
	UserService "go-api/app/user/service"
	"go-api/config"
	"go-api/database"
	"go-api/extparty"
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

	//Init Ext party
	dans := extparty.NewDansConstruct()

	//User Endpoint
	userRepo := UserRepo.NewUserRepository(db)
	userService := UserService.NewUserService(userRepo)
	routes.NewUserHandler(router, userService)

	//Job Endpoint
	jobService := JobService.NewJobService(dans)
	routes.NewJobHandler(router, jobService)

	router.Logger.Fatal(router.Start(":" + conf.GetString("app.port")))
}

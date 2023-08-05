package main

import (
	"echo_sprint_planner/app/domains/services"
	"echo_sprint_planner/app/infra/db"
	"echo_sprint_planner/app/infra/db/repositories"
	"echo_sprint_planner/app/infra/handler"
	"echo_sprint_planner/app/infra/router"
)

func main() {
	db := db.NewDB()
	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	e := router.NewRouter(userHandler)
	e.Logger.Fatal(e.Start(":8080"))
}

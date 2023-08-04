package main

import "echo_sprint_planner/app/infra/router"

func main() {
	e := router.New()
	e.Logger.Fatal(e.Start(":8080"))
}

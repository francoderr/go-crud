package main

import (
	"echo-mongo-api/configs" //add this
	"echo-mongo-api/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	//run database
	configs.ConnectDB()

	//routes
	routes.UserRoute(e) //add this

	e.Logger.Fatal(e.Start(":6000"))
}

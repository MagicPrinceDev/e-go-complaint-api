package main

import (
	"e-complaint-api/config"
	"e-complaint-api/drivers/mysql"
	"e-complaint-api/routes"
	"fmt"

	"github.com/labstack/echo/v4"
)

func main() {
	// config.LoadEnv()
	config.InitConfigMySQL()
	DB := mysql.ConnectDB(config.InitConfigMySQL())
	e := echo.New()

	fmt.Println(DB)

	routes := routes.RouteController{}

	routes.InitRoute(e)
	e.Start(":8080")
}

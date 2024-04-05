package main

import (
	"github.com/labstack/echo/v4"
	"mahasiswa/config"
	"mahasiswa/routes"
)

func main() {
	config.Init()
	defer func() {
		db := config.InitDB()
		dbSQL, _ := db.DB()
		dbSQL.Close()
	}()

	e := echo.New()
	routes.Routes(e)
	e.Logger.Fatal(e.Start(":8000"))
}

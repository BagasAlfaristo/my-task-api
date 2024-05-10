package main

import (
	"my-task-api/app/configs"
	databases "my-task-api/app/database"
	"my-task-api/app/migrations"
	"my-task-api/app/routers"

	"github.com/labstack/echo/v4"
	//	"github.com/labstack/echo/v4"
)

func main() {
	cfg := configs.InitConfig()
	dbMysql := databases.InitDBMysql(cfg)
	migrations.InitMigrations(dbMysql)

	e := echo.New()

	routers.InitRouter(e, dbMysql)
	e.Logger.Fatal(e.Start(":8080"))
}

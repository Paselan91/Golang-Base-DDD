package main

import (
	"github.com/labstack/echo"
	"my-module/app/http/routes"
	"os"
)

func main() {
	e := echo.New()

	routes.Routes(e)
	routes.Run(e, os.Getenv("BE_PORT"))
}

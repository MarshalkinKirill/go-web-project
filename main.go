package main

import (
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	userGroup := e.Group("/user")

}
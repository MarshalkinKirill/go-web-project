package handler

import (
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userController *UserController
}

func NewUserHandler(userController *UserController) *UserHandler {
	return &UserHandler{userController}
}
func InitUserRoutes(g *echo.Group) {
	g.GET("/user", getAllUsers)
	g.POST("/users", createUser)
	g.GET("/users/:id", getUserByID)
	g.PUT("/users/:id", updateUserByID)
	g.DELETE("/users/:id", deleteUserByID)
}

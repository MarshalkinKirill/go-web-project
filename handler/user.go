package handler

import (
	"github.com/labstack/echo/v4"
	"go-web-project/controller"
)

type UserHandler struct {
	userController *controller.UserController
}

func NewUserHandler(userController *controller.UserController) *UserHandler {
	return &UserHandler{userController}
}
func (h *UserHandler) InitUserRoutes(g *echo.Group) {
	g.GET("/user", h.userController.getAllUsers)
	g.POST("/users", h.userController.createUser)
	g.GET("/users/:id", h.userController.getUserByID)
	g.PUT("/users/:id", h.userController.updateUserByID)
	g.DELETE("/users/:id", h.userController.deleteUserByID)
}

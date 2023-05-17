package controller

type UserController struct {
	userRepository UserRpository
}

func NewUserController(r UserRepository) *UserController {
	return &UserController{r}
}

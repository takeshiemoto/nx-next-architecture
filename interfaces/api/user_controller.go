package api

import (
	"toh-echo/domain"
	"toh-echo/interfaces/database"
	"toh-echo/usecase"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	Interactor usecase.UserInteractor
}

func NewUserController(sqlHandler database.SqlHandler) *UserController {
	return &UserController{
		Interactor: usecase.UserInteractor{
			UserRepository: &database.UserRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (ctr *UserController) Create(c echo.Context) {
	u := domain.User{}
	c.Bind(&u)
	ctr.Interactor.Add(u)
	createdUsers := ctr.Interactor.GetInfo()
	c.JSON(201, createdUsers)
}

func (ctr *UserController) GetUser() []domain.User {
	res := ctr.Interactor.GetInfo()
	return res
}

func (ctr *UserController) Delete(id string) {
	ctr.Interactor.Delete(id)
}

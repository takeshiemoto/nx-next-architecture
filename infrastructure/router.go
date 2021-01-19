package infrastructure

import (
	"net/http"
	controllers "toh-echo/interfaces/api"

	"github.com/labstack/echo/v4/middleware"

	"github.com/labstack/echo/v4"
)

func Init() {
	e := echo.New()
	e.Use(middleware.Logger())
	userController := controllers.NewUserController(NewSqlHandler())

	e.GET("/users", func(c echo.Context) error {
		users := userController.GetUser()
		c.Bind(&users)
		return c.JSON(http.StatusOK, users)
	})

	e.POST("/users", func(c echo.Context) error {
		userController.Create(c)
		return c.String(http.StatusOK, "created")
	})

	e.DELETE("/users/:id", func(c echo.Context) error {
		id := c.Param("id")
		userController.Delete(id)
		return c.String(http.StatusOK, "deleted")
	})

	e.Logger.Fatal(e.Start(":1323"))
}

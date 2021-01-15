package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ルーティング
	e.GET("/", index)

	e.Logger.Fatal(e.Start(":1323"))
}

type Ping struct {
	Status int `json:"status"`
}

func index(c echo.Context) error {
	ping := Ping{Status: http.StatusOK}
	return c.JSON(http.StatusOK, ping)
}

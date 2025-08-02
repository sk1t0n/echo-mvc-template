package controllers

import "github.com/labstack/echo/v4"

type Controller interface {
	Index(echo.Context) error
	Create(echo.Context) error
	Store(echo.Context) error
	Show(echo.Context) error
	Edit(echo.Context) error
	Update(echo.Context) error
	Destroy(echo.Context) error
}

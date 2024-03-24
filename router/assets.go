package assets_router

import "github.com/labstack/echo/v4"

func Mount(e *echo.Echo) {
	e.Static("/dist", "dist")
}

package main

import (
	"log"
	"log/slog"

	"github.com/hmbil694/co-author-go/db"
	assets_router "github.com/hmbil694/co-author-go/router"
	"github.com/hmbil694/co-author-go/views/event"
	"github.com/hmbil694/co-author-go/views/home"
	"github.com/hmbil694/co-author-go/views/login"
	"github.com/hmbil694/co-author-go/views/utils"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		slog.Error("Error loading .env file")
		log.Fatal(err)
	}

	db, err := db.NewSQL()
	if err != nil {
		slog.Error("Error loading .env file")
		log.Fatal(err)

	}

	defer db.Close()
	slog.Info("Successfully connected to DB")

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	e.Use(middleware.Recover())

	assets_router.Mount(e)

	e.GET("/", func(c echo.Context) error {
		return utils.Render(c, home.Home())
	})

	e.GET("/login", func(c echo.Context) error {
		return utils.Render(c, login.Login())
	})

	e.GET("/event/:eventId", func(c echo.Context) error {
		return utils.Render(c, event.EventPage())
	})

	e.Logger.Fatal(e.Start(":42069"))
}

package main

import (
	"log/slog"
	"os"
	"strings"

	"github.com/hmbil694/co-author-go/db"
	"github.com/hmbil694/co-author-go/model"
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
		os.Exit(1)
	}

	dbClient, err := db.NewSQL()
	if err != nil {
		slog.Error("Error loading .env file")
		os.Exit(1)

	}

	defer dbClient.Close()
	slog.Info("Successfully connected to DB")

	err = db.CreateTables(dbClient)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	err = db.SeedDataBase(dbClient)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	assets_router.Mount(e)

	e.GET("/", func(c echo.Context) error {
		events := []model.Event{}
		err := dbClient.Select(&events, "SELECT * from event")

		if err != nil {
			slog.Error(err.Error())
			return err
		}

		return utils.Render(c, home.Home(events))
	})

	e.GET("/event", func(c echo.Context) error {
		searchTerm := c.QueryParam("q")

		events := []model.Event{}
		err := dbClient.Select(&events, "SELECT * from event WHERE lower(name) like '%"+strings.ToLower(searchTerm)+"%'")

		if err != nil {
			slog.Error(err.Error())
			return err
		}

		return utils.Render(c, home.EventGroup(events))

	})

	e.GET("/login", func(c echo.Context) error {
		return utils.Render(c, login.Login())
	})

	e.GET("/event/:eventId", func(c echo.Context) error {
		return utils.Render(c, event.EventPage())
	})

	e.Logger.Fatal(e.Start(":42069"))
}

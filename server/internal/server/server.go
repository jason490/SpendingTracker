package server

import (
	"SpendingTracker/internal/storage"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	store *storage.Storage
	e     *echo.Echo
}

func RunServer(store *storage.Storage) *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORS())
	// e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	server := Server{
		store: store,
		e:     e,
	}

	server.routes()
	return e
}

package http

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewHandler() http.Handler {
	log.Println("creating http handler")

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Static("./web/static"))
	e.Use(middleware.Static("./web/template"))
	return e
}

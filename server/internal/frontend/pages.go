package frontend

import (
	"SpendingTracker/internal/templ/pages"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (f *Frontend) getHomePage(c echo.Context) error {
	return Render(c, http.StatusOK, pages.Home())
}

func (f *Frontend) getLoginPage(c echo.Context) error {
	return Render(c, http.StatusOK, pages.Home())
}

func (f *Frontend) getUserHomePage(c echo.Context) error {
	return Render(c, http.StatusOK, pages.UserHomePage())
}

func (f *Frontend) getUserSettings(c echo.Context) error {
	return Render(c, http.StatusOK, pages.Home())
}

package server

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *Server) postCreateUser( c echo.Context ) error {
	fmt.Println(c)
	var test Test = Test{
		Name: "Hello",
		Test: "World!",
	}
	return c.JSON(http.StatusOK, test)
}

func (s *Server) postLogin(c echo.Context) error {
	var test Test = Test{
		Name: "Hello",
		Test: "World!",
	}
	return c.JSON(http.StatusOK, test)

}

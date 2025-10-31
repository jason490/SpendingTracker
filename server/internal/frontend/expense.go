package frontend

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (f *Frontend) postAddExpense( c echo.Context ) error {
	// Set tag
	var test Test = Test{
		Name: "Hello",
		Test: "World!",
	}
	return c.JSON(http.StatusOK, test)
}

func (f *Frontend) postChangeExpense( c echo.Context ) error {
	var test Test = Test{
		Name: "Hello",
		Test: "World!",
	}
	return c.JSON(http.StatusOK, test)
}

func (f *Frontend) postCreateTag( c echo.Context ) error {
	// has is default
	var test Test = Test{
		Name: "Hello",
		Test: "World!",
	}
	return c.JSON(http.StatusOK, test)
}

func (f *Frontend) postChangeTag( c echo.Context ) error {
	// has is default
	var test Test = Test{
		Name: "Hello",
		Test: "World!",
	}
	return c.JSON(http.StatusOK, test)
}

func (f *Frontend) getAllTags( c echo.Context ) error {
	var test Test = Test{
		Name: "Hello",
		Test: "World!",
	}
	return c.JSON(http.StatusOK, test)
}

func (f *Frontend) getTotalExpense( c echo.Context ) error {
	var test Test = Test{
		Name: "Hello",
		Test: "World!",
	}
	return c.JSON(http.StatusOK, test)
}

func (f *Frontend) getAllExpenses( c echo.Context ) error {
	var test Test = Test{
		Name: "Hello",
		Test: "World!",
	}
	return c.JSON(http.StatusOK, test)
}

func (f *Frontend) getExpense( c echo.Context ) error {
	// include tag
	var test Test = Test{
		Name: "Hello",
		Test: "World!",
	}
	return c.JSON(http.StatusOK, test)
}

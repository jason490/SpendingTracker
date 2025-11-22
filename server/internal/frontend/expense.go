package frontend

import (
	"SpendingTracker/internal/storage"
	"SpendingTracker/internal/templ/components"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func (f *Frontend) postAddExpense(c echo.Context) error {
	inter := c.Get("UserData")
	// interface -> storage.User Struct
	user, ok := inter.(storage.User)
	if !ok {
		log.Error("Unable to convert interface to User struct")
	}
	tag := storage.Tag{}
	cost, err := strconv.ParseFloat(c.FormValue("cost"), 32)
	if err != nil {
		log.Error("Unable to parse float")
		return c.String(http.StatusBadRequest, "The cost is not a valid number")
	}
	expense := storage.Expense{
		Name:        c.FormValue("name"),
		Description: "",
		Cost:        cost,
	}
	err = f.store.AddExpense(&expense, &user, &tag)
	if err != nil {
		log.Error("Unable to addExpense")
		return c.String(http.StatusInternalServerError, "Could not add expense. Try again later.")
	}

	return f.getMonthExpense(c)
}

func (f *Frontend) postChangeExpense(c echo.Context) error {
	var test Test = Test{
		Name: "Hello",
		Test: "World!",
	}
	return c.JSON(http.StatusOK, test)
}

func (f *Frontend) postCreateTag(c echo.Context) error {
	// has is default
	var test Test = Test{
		Name: "Hello",
		Test: "World!",
	}
	return c.JSON(http.StatusOK, test)
}

func (f *Frontend) postChangeTag(c echo.Context) error {
	// has is default
	var test Test = Test{
		Name: "Hello",
		Test: "World!",
	}
	return c.JSON(http.StatusOK, test)
}

func (f *Frontend) getAllTags(c echo.Context) error {
	var test Test = Test{
		Name: "Hello",
		Test: "World!",
	}
	return c.JSON(http.StatusOK, test)
}

func (f *Frontend) getTotalExpense(c echo.Context) error {
	var test Test = Test{
		Name: "Hello",
		Test: "World!",
	}
	return c.JSON(http.StatusOK, test)
}

func (f *Frontend) getMonthExpense(c echo.Context) error {
	inter := c.Get("UserData")
	// interface -> storage.User Struct
	user, ok := inter.(storage.User)
	if !ok {
		log.Error("Unable to convert interface to User struct")
	}
	tag := storage.Tag{}
	expenses, err := f.store.GetAllExpenses(&tag, &user)
	fmt.Println((*expenses)[0].CreatedAt)
	if err != nil && err != sql.ErrNoRows {
		log.Error("Unable to get Expenses")
		log.Error(err)
		return Render(c, http.StatusInternalServerError,
			components.ErrorCardMsg("Server error unable to get expenses"))
	}
	return Render(c, http.StatusOK, components.AllExpenses(expenses))
}

func (f *Frontend) getAllExpenses(c echo.Context) error {
	var test Test = Test{
		Name: "Hello",
		Test: "World!",
	}
	return c.JSON(http.StatusOK, test)
}

func (f *Frontend) getExpense(c echo.Context) error {
	// include tag
	var test Test = Test{
		Name: "Hello",
		Test: "World!",
	}
	return c.JSON(http.StatusOK, test)
}

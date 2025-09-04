package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Test struct {
	Name string `json:"name"`
	Test string `json:"test"`
}

func (s *Server) AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// cookie, err := c.Cookie("Tutorfi_Account")
		// if err != nil {
		// 	return c.Redirect(302, "/login")
		// }
		// sessionId := cookie.Value
		// // acc, err := a.store.GetAccountSessionId(sessionId)
		// if err != nil {
		// 	// Oops something happened
		// 	// return some json too
		// 	return c.String(302, "/login")
		// }
		// if acc == nil {
		// 	// Unauthorized
		// 	return c.Redirect(http.StatusUnauthorized, "/login")
		// }
		// return c.String(http.StatusOK,"Hello")
		return next(c)
	}
}

func (s *Server) routes() error {
	var test Test = Test{
		Name: "Hello",
		Test: "World!",
	}
	s.e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, test)
	})
	e := s.e.Group("/api", s.AuthMiddleware)

	e.POST("/post/createUser", s.postCreateUser)
	e.POST("/post/addExpense", s.postAddExpense)
	e.POST("/post/changeExpense", s.postChangeExpense)
	e.POST("/post/changeTag", s.postChangeTag)
	e.POST("/post/createTag", s.postCreateTag)

	e.GET("/get/expense", s.getExpense)
	e.GET("/get/allTags", s.getAllTags)
	e.GET("/get/allExpenses", s.getAllExpenses)
	e.GET("/get/totalExpense", s.getTotalExpense)
	return nil
}

func (s *Server) postCreateUser( c echo.Context ) error {
	var test Test = Test{
		Name: "Hello",
		Test: "World!",
	}
	return c.JSON(http.StatusOK, test)
}

func (s *Server) postAddExpense( c echo.Context ) error {
	// Set tag
	var test Test = Test{
		Name: "Hello",
		Test: "World!",
	}
	return c.JSON(http.StatusOK, test)
}

func (s *Server) postChangeExpense( c echo.Context ) error {
	var test Test = Test{
		Name: "Hello",
		Test: "World!",
	}
	return c.JSON(http.StatusOK, test)
}

func (s *Server) postCreateTag( c echo.Context ) error {
	// has is default
	var test Test = Test{
		Name: "Hello",
		Test: "World!",
	}
	return c.JSON(http.StatusOK, test)
}

func (s *Server) postChangeTag( c echo.Context ) error {
	// has is default
	var test Test = Test{
		Name: "Hello",
		Test: "World!",
	}
	return c.JSON(http.StatusOK, test)
}

func (s *Server) getAllTags( c echo.Context ) error {
	var test Test = Test{
		Name: "Hello",
		Test: "World!",
	}
	return c.JSON(http.StatusOK, test)
}

func (s *Server) getTotalExpense( c echo.Context ) error {
	var test Test = Test{
		Name: "Hello",
		Test: "World!",
	}
	return c.JSON(http.StatusOK, test)
}

func (s *Server) getAllExpenses( c echo.Context ) error {
	var test Test = Test{
		Name: "Hello",
		Test: "World!",
	}
	return c.JSON(http.StatusOK, test)
}

func (s *Server) getExpense( c echo.Context ) error {
	// include tag
	var test Test = Test{
		Name: "Hello",
		Test: "World!",
	}
	return c.JSON(http.StatusOK, test)
}

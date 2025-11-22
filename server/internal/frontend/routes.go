package frontend

import (
	"SpendingTracker/internal/storage"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type Test struct {
	Name string `json:"name"`
	Test string `json:"test"`
}

func createCookie(sessionid string, remember bool) *http.Cookie {
	var cookie = new(http.Cookie)
	cookie.Name = "Account"
	cookie.Value = sessionid
	if remember {
		cookie.Expires = time.Now().AddDate(1, 0, 0)
	} else {
		cookie.Expires = time.Now().AddDate(0, 0, 1)
	}
	cookie.HttpOnly = true
	cookie.Secure = true
	cookie.Path = "/"
	return cookie
}

func (f *Frontend) AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// var test Test = Test{
		// 	Name: "Hello",
		// 	Test: "World!",
		// }
		// c.Set("User", test)
		// val := c.Get("User")
		// Verify api key first cookie, err := c.Cookie("Tutorfi_Account")
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
		// Remove this when done
		user := storage.User {
			Email: "test@test.com",
		}
		
		err := f.store.GetUserFromEmail(&user)
		if err != nil {
			fmt.Println("Unable to get test user")
			log.Fatal(err)
		}
		cookie := createCookie(user.SessionId, true)
		c.SetCookie(cookie)
		c.Set("UserData", user)
		return next(c)
	}
}

func (f *Frontend) routes(){
	e := f.e.Group("/user", f.AuthMiddleware)

	f.e.GET("/", f.getHomePage)

	// Pages
	e.GET("/home", f.getUserHomePage)

	// Components
	e.POST("/post/login", f.postLogin)
	e.POST("/post/createUser", f.postCreateUser)
	e.POST("/post/addExpense", f.postAddExpense)
	e.POST("/post/changeExpense", f.postChangeExpense)
	e.POST("/post/changeTag", f.postChangeTag)
	e.POST("/post/createTag", f.postCreateTag)

	e.GET("/get/expense", f.getExpense)
	e.GET("/get/allTags", f.getAllTags)
	e.GET("/get/allExpenses", f.getAllExpenses)
	e.GET("/get/totalExpense", f.getTotalExpense)
	e.GET("/get/monthExpense", f.getMonthExpense)

	f.e.Static("/css", "./internal/templ/css")
	f.e.Static("/js", "./internal/templ/js")
}


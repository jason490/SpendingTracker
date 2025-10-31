package frontend

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	// "golang.org/x/crypto/bcrypt"
)

func (f *Frontend) postCreateUser( c echo.Context ) error {
	fmt.Println(c)
	var test Test = Test{
		Name: "Hello",
		Test: "World!",
	}
	return c.JSON(http.StatusOK, test)
}

func (f *Frontend) postLogin(c echo.Context) error {
	var test Test = Test{
		Name: "Hello",
		Test: "World!",
	}
	// password := "mySecretPassword123"
	// bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	// hash := string(bytes)
	// if err != nil {
	// 	fmt.Println("Error hashing password:", err)
	// 	return err
	// }
	// err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	// if err != nil {
	// 	fmt.Println("incorrect password", err)
	// 	return err
	// }
	return c.JSON(http.StatusOK, test)
}

package frontend

import (
	"SpendingTracker/internal/storage"
	"database/sql"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
)

type Frontend struct {
	store *storage.Storage
	e     *echo.Echo
}

func RunFrontend(store *storage.Storage, e *echo.Echo){
	// e.Use(middleware.CORS())
	// e.Use(middleware.Recover())
	// e.Use(middleware.Logger())

	user := storage.User {
		Email: "test@test.com",
		Username: "test",
		Password: "test",
	}
	err := store.GetUserFromEmail(&user)
	if err == sql.ErrNoRows {
		fmt.Println("Created Test User")
		err1 := store.CreateUser(&user)
		if err1 != nil {
			fmt.Println("Can not create test user")
			log.Fatal(err1)
		}
		err1 = store.UpdateSessionId(&user)
		fmt.Println(user.SessionId)
		if err1 != nil {
			fmt.Println("Can't update session Id")
			log.Fatal(err1)
		}
		tag := storage.Tag{
			Name: "test",
		}
		expense := storage.Expense{
			Name: "test",
			Cost: 23,
			Description: "",
		}
		err1 = store.CreateTag(&tag, &user)
		if err1 != nil {
			fmt.Println("Can't create Tag")
			log.Fatal(err1)
		}
		err1 = store.AddExpense(&expense, &user, &tag)
		if err1 != nil {
			fmt.Println("Can't create Expense")
			log.Fatal(err1)
		}
	} else if err != nil {
		log.Fatal(err)
	}


	frontend := Frontend{
		store: store,
		e: e,
	}

	frontend.routes()
}

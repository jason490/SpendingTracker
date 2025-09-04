package storage

import (
	"time"

	"github.com/google/uuid"
	"github.com/labstack/gommon/log"
)

type User struct {
	Id uint
	SessionId string
	Username string
	Password string
	Email string
	TotalSpending float64
	Created_at uint
}

func (s *Storage) CreateUser(u *User, x *int) error {

	query := `INSERT INTO "Users"("username","password","email",
			  created_at") VALUES (?,?,?,?); `

	currTime := time.Now().Unix()
	_, err := s.db.Exec(query, u.Username, u.Password, u.Email, u.Password, currTime)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (s *Storage) Login(u *User) error {
	sessionId := uuid.New()
	query := `SELECT "username","password", "id" FROM "Users" WHERE "email" = ?`
	row := s.db.QueryRow(query, u.Email)
	if row.Err() != nil {
		return row.Err()
	}
	err := row.Scan(u.Username, u.Password, u.Id)
	if err != nil {
		return err
	}
	u.SessionId = sessionId.String()
	query = `UPDATE "Users" SET "session_id" = ? WHERE "id" = ?`
	_, err = s.db.Exec(query, sessionId, u.Id)
	if err != nil {
		return err
	}

	return nil
}



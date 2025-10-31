package storage

import (
	"time"

	"github.com/google/uuid"
	"github.com/labstack/gommon/log"
)

type User struct {
	Id            uint
	SessionId     string
	Username      string
	Password      string
	Email         string
	TotalSpending float64
	Created_at    uint
}

// Requires username, password, and email.
// Ensure email is verified unique.
func (s *Storage) CreateUser(u *User) error {

	query := `INSERT INTO "Users"("username","password","email", "created_at", "total_spending")
	VALUES (?, ?, ?, ?, ?); `

	currTime := time.Now().Unix()
	result, err := s.db.Exec(query, u.Username, u.Password, u.Email, currTime, 0.0)
	if err != nil {
		log.Error(err)
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Error(err)
		return err
	}
	u.Id = uint(id)
	return nil
}

// Requires Email
func (s *Storage) GetUserFromEmail(u *User) error {
	query := `SELECT "id","username","session_id","email","total_spending" FROM "Users" 
		WHERE "email" = ?;`
	row := s.db.QueryRow(query, u.Email)
	if row.Err() != nil {
		return row.Err()
	}
	err := row.Scan(&u.Id, &u.Username, &u.SessionId, &u.Email, &u.TotalSpending)
	// Check sql.ErrNoRows
	if err != nil {
		return err
	}
	return nil
}

// Requires Id
func (s *Storage) GetUserFromId(u *User) error {
	query := `SELECT "id","username","session_id","email","total_spending" FROM "Users" 
		WHERE "id" = ?;`

	row := s.db.QueryRow(query, u.Id)
	if row.Err() != nil {
		return row.Err()
	}
	err := row.Scan(&u.Id, &u.Username, &u.SessionId, &u.Email, &u.TotalSpending)
	if err != nil {
		return err
	}
	return nil
}


// Requires User and SessionId
func (s *Storage) GetUserIdFromSessionId(u *User) error {
	query := `SELECT "id" FROM "Users" WHERE "sessionId"=?;`
	row := s.db.QueryRow(query, u.SessionId)

	if row.Err() != nil {
		return row.Err()
	}

	err := row.Scan(u.Id)
	if err != nil {
		return err
	}

	return nil
}

// Requires sessionId and Id
func (s *Storage) UpdateSessionId(u *User) error {
	sessionId := uuid.New()
	u.SessionId = sessionId.String()
	query := `UPDATE "Users" SET "session_id" = ? WHERE "id" = ?;`
	_, err := s.db.Exec(query, u.SessionId, u.Id)
	if err != nil {
		return err
	}
	return nil
}


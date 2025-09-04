package storage

import (
	"database/sql"

	"github.com/labstack/gommon/log"
)

type Storage struct {
	db   *sql.DB
	Data struct {
		Username string
		Expense  float64
		Tag      string
	}
}

func NewSqliteStorage(db *sql.DB) *Storage {
	return &Storage{
		db: db,
	}
}

func (s *Storage) AddExpense(userId string, tagId uint) error {
	query := `INSERT INTO "Expenses"("user_id","tag_id","name","description","cost")
			  VALUES (?, ?, ?, ?, ?); `
	_, err := s.db.Exec(query, s.Data.Username)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (s *Storage) ChangeExpense() error {
	query := `UPDATE "Expenses" SET ("user_id","tag_id","name","description","cost")
			  VALUES (?, ?, ?); `
	_, err := s.db.Exec(query, s.Data.Username)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (s *Storage) CreateTag() error {
	return nil
}

func (s *Storage) ChangeTag() error {
	return nil
}

func (s *Storage) GetAllTags() error {
	return nil
}

func (s *Storage) GetTotalExpenses() error {
	return nil
}

func (s *Storage) GetAllExpenses() error {
	return nil
}

func (s *Storage) GetExpense() error {
	return nil
}

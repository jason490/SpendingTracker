package storage

import (
	"database/sql"
	"errors"
	"strconv"
	"time"

	"github.com/labstack/gommon/log"
)

type Storage struct {
	db *sql.DB
}

type Expense struct {
	Id          int64
	Name        string
	Cost        float64
	Description string
	CreatedAt   int64
}

type Tag struct {
	Id   int64
	Name string
}

type Filters struct {
	StartTime int64
	EndTime   int64
	Tags      *[]Tag
	Limit     int64
	Glob      string
}

func NewSqliteStorage(db *sql.DB) *Storage {
	return &Storage{
		db: db,
	}
}

// Requires all of the values in Expense, user ID, and Tag ID
func (s *Storage) AddExpense(d *Expense, u *User, t *Tag) error {
	query := `INSERT INTO "Expenses"("tag_id","user_id","name","description","cost","created_at")
			  VALUES (?, ?, ?, ?, ?, ?);`
	currTime := time.Now().Unix()
	_, err := s.db.Exec(query, t.Id, u.Id, d.Name, d.Description, d.Cost, currTime)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

// Requires user ID and Name of tag
func (s *Storage) CreateTag(t *Tag, u *User) error {
	query := `INSERT INTO "Tags"("user_id","name") VALUES (?, ?);`
	result, err := s.db.Exec(query, u.Id, t.Name)
	if err != nil {
		log.Error(err)
		return err
	}
	t.Id, err = result.LastInsertId()
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

// Requires all of the values in Expense, user ID, and Tag ID
func (s *Storage) ChangeExpense(d *Expense, t *Tag, u *User) error {
	query := `UPDATE "Expenses" SET "tag_id"=?, "name"=?,"description"=?,"cost"=? 
			  WHERE "user_id"=? AND "id"=?;`
	_, err := s.db.Exec(query, t.Id, d.Name, d.Description, d.Cost, u.Id, d.Id)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

// Requires Tag Name, Tag ID, and user Id
func (s *Storage) ChangeTag(t *Tag, u *User) error {
	query := `UPDATE "Tags" SET "name"=? WHERE "tag_id"=? AND "user_id"=?;`
	_, err := s.db.Exec(query, t.Name, t.Id, u.Id)

	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

// Requires Expense ID
func (s *Storage) DeleteExpense(d *Expense, t *Tag, u *User) error {
	query := ``
	_, err := s.db.Exec(query, t.Id, d.Name, d.Description, d.Cost, u.Id, d.Id)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

// Requires Tag ID
func (s *Storage) DeleteTag(t *Tag, u *User) error {
	query := ``
	_, err := s.db.Exec(query, t.Name, t.Id, u.Id)

	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

// Requires user ID
func (s *Storage) GetAllTags(u *User) (*[]Tag, error) {
	query := `SELECT "id", "name" FROM "Tags" WHERE "user_id"=?;`
	rows, err := s.db.Query(query, u.Id)

	if err != nil {
		log.Error(err)
		return nil, err
	}
	tags := make([]Tag, 0)
	for rows.Next() {
		var tag Tag
		if err := rows.Scan(&tag.Id, &tag.Name); err != nil {
			log.Error(err)
			return nil, err
		}
		tags = append(tags, tag)
	}

	return &tags, nil
}

// Requires user id and tag name
func (s *Storage) GetTagId(t *Tag, u *User) error {
	query := `SELECT "id", "name" FROM "Tags" WHERE "user_id"=? AND "name"=?;`
	row := s.db.QueryRow(query, u.Id, t.Name)

	if row.Err() != nil {
		log.Error(row.Err())
		return row.Err()
	}

	err := row.Scan(&t.Id, &t.Name)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

// Requires user id and tag id
func (s *Storage) GetTagName(t *Tag, u *User) error {
	query := `SELECT "id", "name" FROM "Tags" WHERE "user_id"=? AND "id"=?;`
	row := s.db.QueryRow(query, u.Id, t.Id)

	if row.Err() != nil {
		log.Error(row.Err())
		return row.Err()
	}
	err := row.Scan(&t.Id, &t.Name)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

// Requires UserID
func (s *Storage) GetMonthExpenses(t *Tag, u *User) (*[]Expense, error) {

	year := time.Now().Year()
	month := time.Now().Month()
	firstDayInMonth := time.Date(year, month, 0, 0, 0, 0, 0, time.UTC).Unix()
	lastDayInMonth := time.Date(year, month+1, 0, 23, 59, 59, 0, time.UTC).Unix()

	query := `SELECT "id", "tag_id", "name", "description", "cost", "created_at"
			  FROM "Expenses" WHERE "user_id"=? AND "created_at" BETWEEN ? and ?
			  ORDER BY "created_at";`

	rows, err := s.db.Query(query, u.Id, firstDayInMonth, lastDayInMonth)

	if err != nil {
		log.Error(err)
		return nil, err
	}

	expenses := make([]Expense, 0)
	for rows.Next() {
		var expense Expense
		if err := rows.Scan(&expense.Id, &t.Id, &expense.Name, &expense.Description, &expense.Cost,
			&expense.CreatedAt); err != nil {
			log.Error(err)
			return nil, err
		}
		expenses = append(expenses, expense)
	}
	return &expenses, nil
}

// Requires UserID
func (s *Storage) GetAllExpenses(t *Tag, u *User) (*[]Expense, error) {
	query := `SELECT "id", "tag_id", "name", "description", "cost", "created_at"
			  FROM "Expenses" WHERE "user_id"=?;`
	rows, err := s.db.Query(query, u.Id)

	if err != nil {
		log.Error(err)
		return nil, err
	}
	expenses := make([]Expense, 0)
	for rows.Next() {
		var expense Expense
		if err := rows.Scan(&expense.Id, &t.Id, &expense.Name, &expense.Description, &expense.Cost,
			&expense.CreatedAt); err != nil {
			log.Error(err)
			return nil, err
		}
		expenses = append(expenses, expense)
	}
	return &expenses, nil
}

// Gets all the expenses from the tags (Not finished)
func (s *Storage) GetTagExpenses(d *Expense, t *Tag, u *User) error {
	query := `SELECT ("id", "name", "description", "cost", "created_at")
			  FROM "Expenses" WHERE "user_id"=? AND "tag_id"=?;`
	row := s.db.QueryRow(query, u.Id, t.Id)

	if row.Err() != nil {
		log.Error(row.Err())
		return row.Err()
	}

	err := row.Scan(&d.Id, &d.Name, &d.Description, &d.Cost, &d.CreatedAt)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

// NOT FINISHED
func (s *Storage) GetExpenses(d *Expense, u *User, f *Filters) error {
	query := `SELECT ("id", "name", "description", "cost", "created_at")
			  FROM "Tags" WHERE "user_id"=?`

	if f.StartTime > f.EndTime {
		return errors.New("Invalid Start Time or End Time")
	}

	startTime := strconv.FormatInt(f.StartTime, 10)
	endTime := strconv.FormatInt(f.EndTime, 10)

	if f.StartTime > 0 && f.EndTime > 0 {
		query = query + ` AND "created_at" BETWEEN ` + startTime + " and " +
			endTime
	} else if f.StartTime > 0 {
		query = query + ` AND "created_at" > ` + startTime
	} else if f.EndTime > 0 {
		query = query + ` AND "created_at" < ` + endTime
	}

	if len(*f.Tags) > 0 {
		query = query + " AND "
		for i, tag := range *f.Tags {
			tagId := strconv.FormatInt(tag.Id, 10)
			if i == len(*f.Tags)-1 {
				query = query + `"tag_id"=` + tagId
			} else {
				query = query + `"tag_id"=` + tagId + " or "
			}
		}
	}
	var glob bool = false
	if f.Glob != "" {
		glob = true
		query = query + ` name GLOB ?`
	}

	if f.Limit > 0 {
		query = query + ` LIMIT ` + strconv.FormatInt(f.Limit, 10)
	}

	query = query + ` ORDER BY "created_at";`

	var row *sql.Row
	if glob {
		row = s.db.QueryRow(query, u.Id, f.Glob)
	} else {
		row = s.db.QueryRow(query, u.Id)
	}

	if row.Err() != nil {
		log.Error(row.Err())
		return row.Err()
	}
	err := row.Scan(&d.Id, &d.Name, &d.Description, &d.Cost, &d.CreatedAt)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

// Requires user Id and tag Id
func (s *Storage) GetExpenseName(d *Expense, u *User) error {
	query := `SELECT ("id", "name", "description", "cost")
			  FROM "Tags" WHERE "user_id"=? AND "id"=?;`
	row := s.db.QueryRow(query, u.Id, d.Id)

	if row.Err() != nil {
		log.Error(row.Err())
		return row.Err()
	}
	err := row.Scan(&d.Id, &d.Name, &d.Description, &d.Cost)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

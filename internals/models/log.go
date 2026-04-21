package models

import (
	"database/sql"
	"errors"
	"time"
)

type Log struct {
	ID      int
	Title   string
	Content string
	Tag     string
	Created time.Time
}

type LogModel struct {
	DB *sql.DB
}

func (m *LogModel) Get(id int) (*Log, error) {
	stmt := `SELECT id, title, content, created, tag FROM devboard WHERE id = ?`

	row := m.DB.QueryRow(stmt, id)

	l := &Log{}
	var ErrNoRecord = errors.New("models: no matching record found")
	err := row.Scan(&l.ID, &l.Title, &l.Content, &l.Created, &l.Tag)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		}
		return nil, err
	}

	return l, nil
}
func (m *LogModel) Insert(title, content, tag string) (int, error) {
	stmt := `INSERT INTO logs(title, content, tag, created) VALUES(?, ?, ?, UTC_TIMESTAMP())`

	result, err := m.DB.Exec(stmt, title, content, tag)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}
func (m *LogModel) Latest() ([]*Log, error) {
	stmt := `SELECT id, title, content, tag, created FROM logs ORDER BY created DESC LIMIT 10`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	logs := []*Log{}
	for rows.Next() {
		l := &Log{}
		err := rows.Scan(&l.ID, &l.Title, &l.Content, &l.Tag, &l.Created)
		if err != nil {
			return nil, err
		}
		logs = append(logs, l)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return logs, nil

}

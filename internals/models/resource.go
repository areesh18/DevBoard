package models

import (
	"database/sql"
	"errors"
	"time"
)

type Resource struct {
	ID      int
	Title   string
	URL     string
	Tag     string
	Note    string
	Created time.Time
}
type ResourceModel struct {
	DB *sql.DB
}



func (m *ResourceModel) Get(id int) (*Resource, error) {
	stmt := `SELECT id, title, url, tag, note, created FROM resources WHERE id=?`
	row := m.DB.QueryRow(stmt, id)
	resource := &Resource{}

	err := row.Scan(&resource.ID, &resource.Title, &resource.URL, &resource.Tag, &resource.Note, &resource.Created)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		}
		return nil, err
	}
	return resource, nil
}
func (m *ResourceModel) Insert(title, url, note, tag string) (int, error) {
	stmt := `INSERT INTO resources(title, url, tag, note, created)
			VALUES(?,?,?,?,UTC_TIMESTAMP())`
	result, err := m.DB.Exec(stmt, title, url, tag, note)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}
func (m *ResourceModel) Latest() ([]*Resource, error) {
	stmt := `SELECT id, title, url, tag, note, created FROM resources ORDER BY created DESC LIMIT 10`
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	resources := []*Resource{}
	for rows.Next() {
		resource := &Resource{}
		err := rows.Scan(&resource.ID, &resource.Title, &resource.URL, &resource.Tag, &resource.Note, &resource.Created)
		if err != nil {
			return nil, err
		}
		resources = append(resources, resource)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return resources, nil

}

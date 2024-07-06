package storage

import (
	"database/sql"
	"fmt"
	"search_scraper/src/types"
)

type Storage struct {
	db *sql.DB
}

func Init(db *sql.DB) *Storage {
	return &Storage{db}
}

func (s *Storage) GetList(listType string) ([]types.Link, error) {
	rows, err := s.db.Query("SELECT * FROM " + listType)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var links []types.Link
	for rows.Next() {
		var link types.Link
		if err := rows.Scan(&link.ID, &link.Domain, &link.Url, &link.FilterType); err != nil {
			return nil, err
		}
		links = append(links, link)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return links, nil
}

func (s *Storage) AddLinkToList(listType string, link types.Link) error {
	insertQuery := fmt.Sprintf("INSERT INTO %s (domain, url, filter_type) VALUES (?, ?, ?);", listType)
	_, err := s.db.Exec(insertQuery, link.Domain, link.Url, link.FilterType)
	return err
}

func (s *Storage) GetLinkFromList(listType string, id int) (types.Link, error) {
	insertQuery := fmt.Sprintf("SELECT * FROM %s WHERE id = ?", listType)
	row := s.db.QueryRow(insertQuery, id)

	var link types.Link
	if err := row.Scan(&link.ID, &link.Domain, &link.Url, &link.FilterType); err != nil {
		return types.Link{}, err
	}
	return link, nil
}

func (s *Storage) UpdateLinkInList(listType string, link types.Link) error {
	insertQuery := fmt.Sprintf("UPDATE %s SET domain = ?, url = ?, filter_type = ? WHERE id = ?", listType)
	_, err := s.db.Exec(insertQuery, link.Domain, link.Url, link.FilterType, link.ID)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) DeleteLinkFromList(listType string, id int) error {
	insertQuery := fmt.Sprintf("DELETE FROM %s WHERE id = ?", listType)
	_, err := s.db.Exec(insertQuery, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) ListConteins(listType string, filterType string, link types.Link) bool {
	return false
}

func (s *Storage) ImportFindedlist() {
}

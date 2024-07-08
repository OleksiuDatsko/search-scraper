package storage

import (
	"database/sql"
	"fmt"
	"search_scraper/src/types"
	"search_scraper/src/utils"
)

type Storage struct {
	db *sql.DB
}

var ErrDBIsNil = fmt.Errorf("db is nil")
var ErrOnURL = fmt.Errorf("no url")

func Init(db *sql.DB) (*Storage, error) {
	if db == nil {
		return nil, ErrDBIsNil
	}
	return &Storage{db}, nil
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
	if link.Url == "" {
		return ErrOnURL
	}
	if link.FilterType == "" {
		link.FilterType = "domain"
	}
	if link.Domain == "" {
		link.Domain = utils.GetDomainFromURL(link.Url)
	}

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
	old, err := s.GetLinkFromList(listType, link.ID)
	if err != nil {
		return err
	}
	if link.Url == "" {
		link.Url = old.Url
	}
	if link.Domain == "" {
		link.Domain = old.Domain
	}
	if link.FilterType == "" {
		link.FilterType = old.FilterType
	}

	insertQuery := fmt.Sprintf("UPDATE %s SET domain = ?, url = ?, filter_type = ? WHERE id = ?", listType)
	_, err = s.db.Exec(insertQuery, link.Domain, link.Url, link.FilterType, link.ID)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) DeleteLinkFromList(listType string, id int) error {
	_, err := s.GetLinkFromList(listType, id)
	if err != nil {
		return err
	}
	insertQuery := fmt.Sprintf("DELETE FROM %s WHERE id = ?", listType)
	_, err = s.db.Exec(insertQuery, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) CleanList(listType string) error {
	insertQuery := fmt.Sprintf("DELETE FROM %s", listType)
	_, err := s.db.Exec(insertQuery)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) ConteinsLinkInList(listType string, link types.Link) (bool, error) {
	_, err := s.db.Query("SELECT * FROM " + listType + " WHERE domain = '" + link.Domain + "'")
	if err != nil {
		return false, err
	}
	return true, nil
}

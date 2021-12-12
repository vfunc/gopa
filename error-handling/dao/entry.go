package dao

import (
	"github.com/pkg/errors"
	"gopa/model"
)

func GetEntryById(id int) (*model.Entry, error) {
	var entry model.Entry
	db := GetDB()
	row := db.QueryRow("select * from entries where id=?", id)
	err := row.Scan(&entry.ID, &entry.Title, &entry.Content, &entry.CreatedAt, &entry.UpdatedAt)
	if err != nil {
		return nil, errors.Wrapf(err, "GetEntryById: %d", id)
	}
	return &entry, nil
}

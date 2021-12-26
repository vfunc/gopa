package biz

import "gopa/internal/data"

type EntryBiz struct {
	repo *data.EntryRepo
}

func NewEntryBiz(repo *data.EntryRepo) *EntryBiz {
	return &EntryBiz{
		repo: repo,
	}
}

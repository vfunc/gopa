package service

import "gopa/internal/biz"

type EntryService struct {
	biz *biz.EntryBiz
}

func NewEntryService(biz *biz.EntryBiz) *EntryService {
	return &EntryService{
		biz: biz,
	}
}

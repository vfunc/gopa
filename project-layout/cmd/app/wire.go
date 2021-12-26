package main

import (
	"github.com/google/wire"
	"gopa/internal/biz"
	"gopa/internal/data"
	"gopa/internal/service"
)

func InitEntryService() *service.EntryService {
	wire.Build(service.NewEntryService, biz.NewEntryBiz, data.NewEntryRepo, data.NewDB)
	return &service.EntryService{}
}

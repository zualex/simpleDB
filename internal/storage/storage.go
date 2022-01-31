package storage

import (
	"github.com/zualex/simpledb/internal/storage/csv"
	storage_interface "github.com/zualex/simpledb/internal/storage/interface"
)

func getProvider(name string) storage_interface.Storage {
	return csv.NewCsv(name)
}

func Create(name string, fields []string) error {
	return getProvider(name).Create(fields)
}

package storage

import (
	"github.com/zualex/simpledb/internal/storage/csv"
)

func Create(name string, fields []string) error {
	return csv.Create(name, fields)
}

package csv

import (
	"encoding/csv"
	"errors"
	"os"

	"github.com/zualex/simpledb/internal/file"
	storage_interface "github.com/zualex/simpledb/internal/storage/interface"
)

type csvStorage struct {
	name string
}

func NewCsv(name string) storage_interface.Storage {
	return &csvStorage{name}
}

func (csvStorage *csvStorage) Create(fields []string) error {
	name := csvStorage.name
	tableName := csvStorage.GetInternalName(name)
	if file.Exists(tableName) {
		return errors.New("Таблица " + name + " уже существует")
	}

	csvFile, err := os.Create(tableName)
	if err != nil {
		return err
	}

	csvwriter := csv.NewWriter(csvFile)
	csvwriter.Write(fields)

	csvwriter.Flush()
	csvFile.Close()

	return nil
}

func (csvStorage *csvStorage) GetInternalName(name string) string {
	return name + ".csv"
}

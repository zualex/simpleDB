package csv

import (
	"bufio"
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
	tableName := csvStorage.GetInternalName()
	if file.Exists(tableName) {
		return errors.New("Таблица " + name + " уже существует")
	}

	csvFile, err := file.Create(tableName)
	if err != nil {
		return err
	}

	csvwriter := csv.NewWriter(csvFile)
	csvwriter.Write(fields)

	csvwriter.Flush()
	csvFile.Close()

	return nil
}

func (csvStorage *csvStorage) GetInternalName() string {
	return csvStorage.name + ".csv"
}

func (csvStorage *csvStorage) GetScheme() ([]string, error) {
	file, err := openFile(csvStorage.GetInternalName())
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	fieldString := scanner.Text() // TODO распарсить на []string

	file.Close()

	return []string{fieldString}, nil
}

func openFile(filename string) (*os.File, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, errors.New("Не удалось открыть " + filename + " ошибка: " + err.Error())
	}

	return file, nil
}

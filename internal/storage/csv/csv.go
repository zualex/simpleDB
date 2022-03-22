package csv

import (
	"bufio"
	"encoding/csv"
	"errors"
	"os"

	"github.com/zualex/simpledb/internal/file"
	parser_field "github.com/zualex/simpledb/internal/parser/field"
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

func (csvStorage *csvStorage) Insert(valueMap map[string]string) error {
	name := csvStorage.name
	tableName := csvStorage.GetInternalName()
	if file.Exists(tableName) == false {
		return errors.New("Таблица " + name + " не существует")
	}

	csvFile, err := file.Open(tableName)
	if err != nil {
		return err
	}

	csvreader := csv.NewReader(csvFile)
	headerFields, err := csvreader.Read()

	values := []string{}
	for _, header := range headerFields {
		if val, ok := valueMap[header]; ok {
			values = append(values, val)
		} else {
			values = append(values, "") // TODO прокидывать более умное дефолтное значение
		}
	}

	csvwriter := csv.NewWriter(csvFile)
	csvwriter.Write(values)

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
	fieldString := scanner.Text()

	file.Close()

	return parser_field.GetFieldsFromString(fieldString), nil
}

func openFile(filename string) (*os.File, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, errors.New("Не удалось открыть " + filename + " ошибка: " + err.Error())
	}

	return file, nil
}

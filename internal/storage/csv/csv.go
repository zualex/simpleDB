package csv

import (
	"encoding/csv"
	"os"
)

func Create(name string, fields []string) error {
	csvFile, err := os.Create(name + ".csv")

	if err != nil {
		return err
	}

	csvwriter := csv.NewWriter(csvFile)
	csvwriter.Write(fields)

	csvwriter.Flush()
	csvFile.Close()

	return nil
}

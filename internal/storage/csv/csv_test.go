package csv

import (
	"os"
	"testing"

	"github.com/zualex/simpledb/internal/file"
	storage_interface "github.com/zualex/simpledb/internal/storage/interface"
)

var name string = "_test"
var csvPackageTest storage_interface.Storage = NewCsv(name)

func TestMain(m *testing.M) {
	internalName := csvPackageTest.GetInternalName()
	file.RemoveIfExists(internalName)

	code := m.Run()

	file.RemoveIfExists(internalName)

	os.Exit(code)
}

func TestCreateSuccess(t *testing.T) {
	fields := []string{"id", "field"}
	err := csvPackageTest.Create(fields)
	if err != nil {
		t.Errorf("storage.Create(\"%s\") error: %s", name, err)
	}

	scheme, err := csvPackageTest.GetScheme()
	if len(scheme) != 2 {
		t.Errorf("csvPackageTest.GetScheme() length: %d", len(scheme))
	}
	if scheme[0] != fields[0] {
		t.Errorf("csvPackageTest.GetScheme() scheme[0]: %s, expected: %s", scheme[0], fields[0])
	}
	if scheme[1] != fields[1] {
		t.Errorf("csvPackageTest.GetScheme() scheme[1]: %s, expected: %s", scheme[1], fields[1])
	}
}

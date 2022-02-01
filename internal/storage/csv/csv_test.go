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
	err := csvPackageTest.Create([]string{"id", "field"})
	if err != nil {
		t.Errorf("storage.Create(\"%s\") error: %s", name, err)
	}

	scheme, err := csvPackageTest.GetScheme()
	t.Error(scheme)
	t.Error(err)
}

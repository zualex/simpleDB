package storage_interface

type Storage interface {
	Create(fields []string) error
	Insert(valueMap map[string]string) error
	GetInternalName() string
	GetScheme() ([]string, error)
}

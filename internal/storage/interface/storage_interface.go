package storage_interface

type Storage interface {
	Create(fields []string) error
	GetInternalName() string
	GetScheme() ([]string, error)
}

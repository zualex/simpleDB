package file

import "os"

func Create(filename string) (*os.File, error) {
	return os.Create(filename)
}

func Exists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}

func Remove(filename string) error {
	return os.Remove(filename)
}

func RemoveIfExists(filename string) error {
	if Exists(filename) {
		return Remove(filename)
	}

	return nil
}

package helper

import (
	"encoding/json"
	"io"
	"os"
)

var (
	ConfigFile     *DataFile
	TasksFile      *DataFile
	CharactersFile *DataFile
	ProductsFile   *DataFile
)

type DataFile struct {
	*os.File
}

func InitFile(path string) (*DataFile, error) {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	}
	return &DataFile{f}, nil
}

func (d *DataFile) GetBytes() ([]byte, error) {
	return io.ReadAll(d.File)
}

func (d *DataFile) ToObject(obj interface{}) error {
	bytes, err := io.ReadAll(d.File)
	if err != nil {
		return err
	}
	if len(bytes) == 0 {
		return nil
	}
	return json.Unmarshal(bytes, obj)
}

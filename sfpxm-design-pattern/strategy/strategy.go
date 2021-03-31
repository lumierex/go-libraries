package strategy

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"os"
)

// 策略模式

type StorageStrategy interface {
	save(name string, data []byte) error
}

var fileSaver = map[string]StorageStrategy{
	//"file
	"file":         &fileStorage{},
	"encrypt_file": &encryptFileStorage{},
}

func NewStorageStrategy(t string) (StorageStrategy, error) {
	f, ok := fileSaver[t]
	if !ok {
		return nil, fmt.Errorf("not found storageStrategy")
	}
	return f, nil
}

// save to file

type fileStorage struct{}

func (f *fileStorage) save(name string, data []byte) error {
	return ioutil.WriteFile(name, data, os.ModeAppend)
}

// encrypt save

type encryptFileStorage struct{}

func (f *encryptFileStorage) save(name string, src []byte) error {
	data, err := encrypt(src)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(name, data, os.ModeAppend)
}

func encrypt(src []byte) ([]byte, error) {
	hash := md5.New()
	sum := hash.Sum(src)
	return sum, nil
}

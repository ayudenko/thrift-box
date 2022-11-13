package main

import (
	"os"
)

type ThriftBox struct {
	storage string
}

type Storage struct {
	fileName string
}

func isError(err error) bool {
	return (err != nil)
}

func (*Storage) load(storageDirPath, storageFileName string) (*os.File, error) {
	var file os.File
	_, err := os.Stat(storageDirPath + storageFileName)
	if isError(err) {
		_, err = os.Stat(storageDirPath)
	} else {
		return os.OpenFile(storageDirPath+storageFileName, os.O_RDWR, 0644)
	}
	if isError(err) {
		err = os.MkdirAll(STORAGE_DIR_PATH, 0755)
		if isError(err) {
			panic(err.Error())
		}
		return os.OpenFile(storageDirPath+storageFileName, os.O_RDWR|os.O_CREATE, 0644)
	}

	return &file, err
}

const STORAGE_DIR_PATH = "/Users/ayudenko/thrift-box/"
const STORAGE_FILE_NAME = "data.bin"

func main() {
	//arg := os.Args[1]
	//box := ThriftBox{storage: "111"}
	storage := Storage{}
	_, err := storage.load(STORAGE_DIR_PATH, STORAGE_FILE_NAME)
	if isError(err) {
		panic(err)
	}
}

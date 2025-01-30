package filestorage

import "sync"

type FileStorage struct {
}

var FileStorageInstanceOnce sync.Once
var FileStorageInstance *FileStorage

func GetFileStorageInstance() *FileStorage {
	FileStorageInstanceOnce.Do(func() {
		FileStorageInstance = &FileStorage{}
	})
	return FileStorageInstance
}

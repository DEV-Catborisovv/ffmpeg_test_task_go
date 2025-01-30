package filestorage

import (
	"fmt"
	"os"
	"path/filepath"
)

const FileStoragePath = "./file_storage/"

func (s *FileStorage) SaveFile(file_data []byte, filename string) error {
	if err := os.MkdirAll(FileStoragePath, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	filePath := filepath.Join(FileStoragePath, filename)

	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	if _, err := file.Write(file_data); err != nil {
		return fmt.Errorf("failed to write data to file: %w", err)
	}

	return nil
}

func (s *FileStorage) IsExistingFile(filename string) bool {
	filePath := filepath.Join(FileStoragePath, filename)
	_, error := os.Stat(filePath)

	if os.IsNotExist(error) {
		return false
	} else {
		return true
	}
}

func (s *FileStorage) DeleteFile(filename string) error {
	filePath := filepath.Join(FileStoragePath, filename)

	if !s.IsExistingFile(filename) {
		return fmt.Errorf("File %s is not existed", filename)
	}

	err := os.Remove(filePath)
	if err != nil {
		return fmt.Errorf("Can't delete file %s: %v", filename, err)
	}

	return nil
}

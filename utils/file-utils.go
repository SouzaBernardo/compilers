package utils

import "os"

func ReadFile(filename *string) (string, error) {
	data, err := os.ReadFile(*filename)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func CreateFileIfNotExists(filename *string) (*os.File, error) {
	file, err := os.OpenFile(*filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}
	return file, nil
}
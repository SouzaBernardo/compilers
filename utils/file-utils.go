package utils

import "os"

func ReadFile(filename *string) (string, error) {
	file, err := os.ReadFile(*filename)
	if err != nil {
		return "", err
	}
	return string(file), nil
}

func OpenFile(filename *string) (*os.File, error) {
	file, err := os.OpenFile(*filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}
	//defer file.Close()
	return file, nil
}

func WriteFile(file *os.File, content *string) error {
	_, err := file.WriteString(*content)
	if err != nil {
		return err
	}
	return nil
}

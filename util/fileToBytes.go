package util

import (
	"io"
	"mime/multipart"
)

func FileToBytes(file multipart.File) ([]byte, error) {
	defer file.Close()

	// Lê todo o conteúdo do arquivo para um slice de bytes
	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return data, nil
}

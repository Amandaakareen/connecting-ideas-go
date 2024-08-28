package repository

import (
	"bytes"
	"context"
	"fmt"

	"github.com/example/infra"
	"github.com/minio/minio-go/v7"
)

func SavePhoto(objectName, filePath string, photo []byte) (string, error) {

	fileBuffer := bytes.NewReader(photo)
	fileSize := int64(fileBuffer.Len())
	fmt.Println("aqui 3")

	uploadInfo, err := infra.Minio.PutObject(context.Background(), "dados", objectName, fileBuffer, fileSize, minio.PutObjectOptions{ContentType: "image/jpeg"})
	if err != nil {
		return "", err
	}

	return uploadInfo.Key, nil
}

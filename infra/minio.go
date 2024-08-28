package infra

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var Minio *minio.Client

func ConnectMinio() {
	minioClient, err := minio.New("localhost:9000", &minio.Options{
		Creds:  credentials.NewStaticV4("guest", "guest", ""),
		Secure: false,
	})

	if err != nil {
		panic(err)
	}

	Minio = minioClient
}

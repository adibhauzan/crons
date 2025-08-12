package configs

import (
	"log"

	"github.com/minio/minio-go/v7"

	"github.com/minio/minio-go/v7/pkg/credentials"
)

func NewMinioConnection() (*minio.Client, error) {
	client, err := minio.New((MinioURL), &minio.Options{
		Creds:  credentials.NewStaticV4(MinioAccessKey, MinioSecretKey, ""),
		Secure: false,
	})

	if err != nil {
		log.Println("Error connect minio: ", err)
		return nil, err
	}

	return client, nil
}

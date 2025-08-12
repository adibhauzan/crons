package repositories

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"image"
	"io"
	"log"
	"mime"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/adibhauzan/crons/configs"
	"github.com/disintegration/imaging"
	"github.com/minio/minio-go/v7"
)

type StorageRepository interface {
	UploadFileMinio(fileName string, buffer multipart.File, contentType string, fileSize int64) (string, error)
	UploadRawFileMinio(fileName string, reader io.Reader, contentType string, fileSize int64) (string, error)
	GetFileBase64(fileName string) (string, error)
	GetResizedFileBase64(filePaths string, maxWidth int) (string, error)
	GetFileMinio(filePaths string) (string, error)
}

type storageRepository struct {
	minio *minio.Client
}

func NewStorageRepository(upload *minio.Client) *storageRepository {
	return &storageRepository{upload}
}

func (repository *storageRepository) UploadFileMinio(fileName string, buffer multipart.File, contentType string, fileSize int64) (string, error) {
	var bucket = configs.MinioBucketName
	_, err := repository.minio.PutObject(context.Background(), bucket, fileName, buffer, fileSize, minio.PutObjectOptions{
		ContentType: contentType,
	})

	if err != nil {
		log.Println("Error uploading file: ", err)
		return "", err
	}
	fileURL := fmt.Sprintf("/%s", fileName)
	fmt.Println(fileURL)
	return fileURL, nil
}

func (repository *storageRepository) UploadRawFileMinio(fileName string, reader io.Reader, contentType string, fileSize int64) (string, error) {
	var bucket = configs.MinioBucketName

	_, err := repository.minio.PutObject(context.Background(), bucket, fileName, reader, fileSize, minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		log.Println("Error uploading raw file: ", err)
		return "", err
	}

	fileURL := fmt.Sprintf("/%s", fileName)
	fmt.Println(fileURL)
	return fileURL, nil
}

func (repository *storageRepository) GetFileMinio(fileName string) (string, error) {
	var bucket = configs.MinioBucketName

	fileUrl, err := repository.minio.PresignedGetObject(context.Background(), bucket, fileName, 3*time.Minute, nil)
	if err != nil {
		log.Println("Error presigned URL: ", err)
		return "", err
	}

	return fileUrl.String(), nil
}

func (repository *storageRepository) GetFileBase64(filePaths string) (string, error) {
	var bucket = configs.MinioBucketName

	fileName := strings.TrimPrefix(filePaths, "/")

	object, err := repository.minio.GetObject(context.Background(), bucket, fileName, minio.GetObjectOptions{})
	if err != nil {
		log.Println("Error get object: ", err)
		return "", err
	}
	defer object.Close()

	head := make([]byte, 512)
	n, err := object.Read(head)
	if err != nil && err != io.EOF {
		log.Println("Error reading object header: ", err)
		return "", err
	}

	contentType := http.DetectContentType(head[:n])
	if contentType == "application/octet-stream" {
		ext := filepath.Ext(fileName)
		contentType = mime.TypeByExtension(ext)
		if contentType == "" {
			contentType = "application/octet-stream"
		}
	}

	fileBytes := append([]byte{}, head[:n]...)
	rest, err := io.ReadAll(object)
	if err != nil {
		log.Println("Error reading full object: ", err)
		return "", err
	}
	fileBytes = append(fileBytes, rest...)

	base64Str := base64.StdEncoding.EncodeToString(fileBytes)

	dataURI := fmt.Sprintf("data:%s;base64,%s", contentType, base64Str)
	return dataURI, nil
}

func (repository *storageRepository) GetResizedFileBase64(filePaths string, maxWidth int) (string, error) {
	var bucket = configs.MinioBucketName
	objectName := strings.TrimPrefix(filePaths, "/")

	obj, err := repository.minio.GetObject(context.Background(), bucket, objectName, minio.GetObjectOptions{})
	if err != nil {
		log.Println("GetObject error:", err)
		return "", err
	}
	defer obj.Close()

	img, format, err := image.Decode(obj)
	if err != nil {
		log.Println("Image decode error:", err)
		return "", err
	}

	resized := imaging.Resize(img, maxWidth, 0, imaging.Lanczos)

	var buf bytes.Buffer
	switch format {
	case "jpeg":
		err = imaging.Encode(&buf, resized, imaging.JPEG)
	case "png":
		err = imaging.Encode(&buf, resized, imaging.PNG)
	default:
		err = imaging.Encode(&buf, resized, imaging.PNG)
	}
	if err != nil {
		log.Println("Image encode error:", err)
		return "", err
	}

	contentType := "image/png"
	if format == "jpeg" {
		contentType = "image/jpeg"
	}

	base64Str := base64.StdEncoding.EncodeToString(buf.Bytes())
	dataURI := fmt.Sprintf(`src="data:%s;base64,%s"`, contentType, base64Str)
	return dataURI, nil
}

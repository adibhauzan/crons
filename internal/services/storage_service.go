package services

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"github.com/adibhauzan/crons/internal/domain/dto"
	"github.com/adibhauzan/crons/internal/repositories"
)

type StorageService interface {
	UploadFile(file dto.UploadFile) (string, error)
	GetFile(fileName string) (string, error)
	GetFileBase64(filePaths string) (string, error)
	GetResizedFileBase64(filePaths string, maxWidth int) (string, error)
	UploadRawFile(file dto.UploadRawFile) (string, error)
}

type storageService struct {
	repo repositories.StorageRepository
}

func NewStorageService(repo repositories.StorageRepository) *storageService {
	return &storageService{repo}
}

func (u *storageService) UploadFile(file dto.UploadFile) (string, error) {

	buffer, err := file.File.Open()
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}
	defer buffer.Close()

	fileBuffer := buffer
	contentType := file.File.Header.Get("Content-Type")
	fileSize := file.File.Size
	folderName := file.FolderName

	originalFileName := file.File.Filename

	// Ekstensi file
	ext := strings.ToLower(filepath.Ext(originalFileName))
	if ext == "" {
		return "", errors.New("file has no extension")
	}

	allowedTYPES := map[string]string{
		"application/pdf":    "pdf",
		"application/msword": "doc",
		"application/vnd.openxmlformats-officedocument.wordprocessingml.document": "docx",
		"application/vnd.ms-excel": "xls",
		"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet":         "xlsx",
		"application/vnd.ms-powerpoint":                                             "ppt",
		"application/vnd.openxmlformats-officedocument.presentationml.presentation": "pptx",
		"text/plain":      "txt",
		"application/rtf": "rtf",
		"text/csv":        "csv",

		// Images
		"image/jpg":     "jpg",
		"image/jpeg":    "jpeg",
		"image/png":     "png",
		"image/gif":     "gif",
		"image/bmp":     "bmp",
		"image/tiff":    "tiff",
		"image/webp":    "webp",
		"image/svg+xml": "svg",
	}

	fileExt, isValidType := allowedTYPES[contentType]
	if !isValidType {
		return "", errors.New("file type is not allowed")
	}

	hash := sha256.Sum256([]byte(originalFileName + time.Now().String()))
	hashedFileName := hex.EncodeToString(hash[:])

	fmt.Println(folderName)

	fileName := fmt.Sprintf("%s/%s.%s", folderName, hashedFileName, fileExt)

	filePath, err := u.repo.UploadFileMinio(fileName, fileBuffer, contentType, fileSize)
	if err != nil {
		return "", fmt.Errorf("failed to Storage file to MinIO: %w", err)
	}

	return filePath, nil
}

func (u *storageService) GetFile(fileName string) (string, error) {

	fileUrl, err := u.repo.GetFileMinio(fileName)
	if err != nil {
		return "", err
	}

	return fileUrl, nil
}

func (u *storageService) UploadRawFile(file dto.UploadRawFile) (string, error) {
	if file.FileName == "" {
		return "", errors.New("filename is required")
	}

	if file.ContentType == "" {
		return "", errors.New("content type is required")
	}

	ext := strings.ToLower(filepath.Ext(file.FileName))
	if ext == "" {
		return "", errors.New("file has no extension")
	}
	allowedTYPES := map[string]string{
		"application/pdf":    "pdf",
		"application/msword": "doc",
		"application/vnd.openxmlformats-officedocument.wordprocessingml.document": "docx",
		"application/vnd.ms-excel": "xls",
		"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet":         "xlsx",
		"application/vnd.ms-powerpoint":                                             "ppt",
		"application/vnd.openxmlformats-officedocument.presentationml.presentation": "pptx",
		"text/plain":      "txt",
		"application/rtf": "rtf",
		"text/csv":        "csv",

		// Images
		"image/jpg":     "jpg",
		"image/jpeg":    "jpeg",
		"image/png":     "png",
		"image/gif":     "gif",
		"image/bmp":     "bmp",
		"image/tiff":    "tiff",
		"image/webp":    "webp",
		"image/svg+xml": "svg",
	}

	_, isValidType := allowedTYPES[file.ContentType]
	if !isValidType {
		return "", fmt.Errorf("file type '%s' is not allowed", file.ContentType)
	}

	hash := sha256.Sum256([]byte(file.FileName + time.Now().String()))
	hashedFileName := hex.EncodeToString(hash[:])
	storagePath := fmt.Sprintf("%s/%s%s", file.FolderName, hashedFileName, ext)

	reader := strings.NewReader(string(file.Content))
	filePath, err := u.repo.UploadRawFileMinio(storagePath, reader, file.ContentType, int64(len(file.Content)))
	if err != nil {
		return "", fmt.Errorf("failed to upload raw file to MinIO: %w", err)
	}

	return filePath, nil
}

func (u *storageService) GetFileBase64(filePaths string) (string, error) {
	return u.repo.GetFileBase64(filePaths)
}

func (u *storageService) GetResizedFileBase64(filePaths string, maxWidth int) (string, error) {
	return u.repo.GetResizedFileBase64(filePaths, maxWidth)
}

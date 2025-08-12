package dto

import "mime/multipart"

type UploadFile struct {
	FolderName string                `form:"folder_name" binding:"required"`
	File       *multipart.FileHeader `form:"file"`
}

type UploadRawFile struct {
	FolderName  string
	FileName    string
	ContentType string
	Content     []byte
}

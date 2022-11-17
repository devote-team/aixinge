package oss

import (
	"mime/multipart"
)

type OSS interface {
	UploadFile(file *multipart.FileHeader) (string, string, string, error)

	FGetObject(key, infile string) error

	GetObject(key string) ([]byte, error)

	DeleteFile(key string) error
}

func NewLocal() OSS {
	return &Local{}
}

package oss

import (
	"aixinge/global"
	"aixinge/utils"
	"errors"
	"io"
	"mime/multipart"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

type Local struct{}

// UploadFile 上传文件
func (*Local) UploadFile(file *multipart.FileHeader) (string, string, string, error) {
	// 读取文件后缀
	ext := path.Ext(file.Filename)
	// 读取文件名并加密
	name := strings.TrimSuffix(file.Filename, ext)
	name = utils.GetByteMd5([]byte(name))
	yearMonth := time.Now().Format("200601")
	// 尝试创建此路径
	var uploadPath = global.CONFIG.Upload.Path
	mkdirErr := os.MkdirAll(filepath.Join(uploadPath, yearMonth), os.ModePerm)
	if mkdirErr != nil {
		return "", "", ext, errors.New("function os.MkdirAll() Filed, err:" + mkdirErr.Error())
	}
	// 拼接新文件名
	filename := name + time.Now().Format("02150405") + ext
	// 路径/年月/文件名
	p := filepath.Join(uploadPath, yearMonth, filename)

	f, openError := file.Open() // 读取文件
	if openError != nil {
		return "", "", ext, errors.New("function file.Open() Filed, err:" + openError.Error())
	}
	defer func(f multipart.File) {
		_ = f.Close()
	}(f) // 创建文件 defer 关闭

	out, createErr := os.Create(p)
	if createErr != nil {
		return "", "", ext, errors.New("function os.Create() Filed, err:" + createErr.Error())
	}
	defer func(out *os.File) {
		_ = out.Close()
	}(out) // 创建文件 defer 关闭

	_, copyErr := io.Copy(out, f) // 传输（拷贝）文件
	if copyErr != nil {
		return "", "", ext, errors.New("function io.Copy() Filed, err:" + copyErr.Error())
	}
	return p, utils.GetFileMd5(p), ext, nil
}

func (*Local) FGetObject(key, infile string) error {
	return nil
}

func (*Local) GetObject(key string) ([]byte, error) {
	return os.ReadFile(key)
}

// DeleteFile 删除文件
func (*Local) DeleteFile(key string) error {
	var uploadPath = global.CONFIG.Upload.Path
	p := filepath.Join(uploadPath, key)
	if strings.Contains(p, uploadPath) {
		if err := os.Remove(p); err != nil {
			return errors.New("本地文件删除失败, err:" + err.Error())
		}
	}
	return nil
}

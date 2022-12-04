package system

import (
	"aixinge/api/model/system"
	"aixinge/global"
	"aixinge/utils"
	"time"
)

type FileService struct {
}

func (f *FileService) GetById(id int64) (err error, file system.File) {
	err = global.DB.First(&file, "id = ?", id).Error
	return err, file
}

func (f *FileService) GetByMd5(md5 string) (err error, file system.File) {
	err = global.DB.First(&file, "md5 = ?", md5).Error
	return err, file
}

func (f *FileService) Save(md5, path, ext, contentType, filename string, size int64) (error, system.File) {
	if len(filename) > 255 {
		filename = filename[0:253]
	}
	file := system.File{ID: utils.Id(), CreatedAt: time.Now(), Md5: md5, Path: path, Ext: ext, Filename: filename,
		ContentType: contentType, Size: size}
	err := global.DB.Create(&file).Error
	// 不显示存储路径
	file.Path = ""
	return err, file
}

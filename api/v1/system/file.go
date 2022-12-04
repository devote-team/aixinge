package system

import (
	"aixinge/api/model/common/response"
	"aixinge/utils"
	"aixinge/utils/oss"
	"errors"
	"github.com/gofiber/fiber/v2"
	"io"
	"net/http"
	"strconv"
)

type File struct{}

var ossClient = oss.NewLocal()

// Download
// @Tags File
// @Summary 文件下载
// @Security ApiKeyAuth
// @Accept  x-www-form-urlencoded
// @Produce application/octet-stream
// @Param id query uint64 true  "主键"
// @Success 200
// @Router /v1/file/download [get]
func (f *File) Download(c *fiber.Ctx) error {
	var id = c.Query("id")
	if len(id) == 0 {
		return errors.New("query param `id` not found")
	}

	id64, err := strconv.ParseInt(id, 10, 64)
	err, file := fileService.GetById(id64)
	if err != nil {
		return errors.New("not found file")
	}
	// 读取文件
	fileBytes, err := ossClient.GetObject(file.Path)
	if err != nil {
		return errors.New("read file error")
	}

	// 附件下载
	c.Attachment(file.Filename)
	return c.Send(fileBytes)
}

// Upload
// @Tags File
// @Summary 文件上传
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param file formData file true "上传文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"上传成功"}"
// @Router /v1/file/upload [post]
func (f *File) Upload(c *fiber.Ctx) error {
	header, err := c.FormFile("file")
	if err != nil {
		return response.FailWithMessage("上传文件不存在", c)
	}

	// 读取文件
	file, err := header.Open()

	// 读取文件字节信息
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return response.FailWithMessage(err.Error(), c)
	}
	contentType := http.DetectContentType(fileBytes)

	// 文件 MD5 值获取判断是否上传
	var fileMd5 = utils.GetByteMd5(fileBytes)
	err, dbFile := fileService.GetByMd5(fileMd5)
	if err == nil {
		return response.OkWithData(dbFile.ID, c)
	}

	// 存储文件
	filePath, md5, ext, uploadErr := ossClient.UploadFile(header)
	if uploadErr != nil {
		return response.FailWithMessage(err.Error(), c)
	}

	err, _file := fileService.Save(md5, filePath, ext, contentType, header.Filename, header.Size)
	if err != nil {
		return response.FailWithMessage(err.Error(), c)
	}
	return response.OkWithData(_file, c)
}

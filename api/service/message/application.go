package message

import (
	"aixinge/api/model/message"
	"aixinge/global"
	"aixinge/utils"
)

type ApplicationService struct {
}

func (c *ApplicationService) Create(app message.Application) (err error) {
	app.ID = utils.Id()
	// 状态，1、正常 2、禁用
	app.Status = 1
	err = global.DB.Create(&app).Error
	return err
}

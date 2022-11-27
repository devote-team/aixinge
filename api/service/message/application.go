package message

import (
	"aixinge/api/model/common/request"
	"aixinge/api/model/message"
	"aixinge/global"
	"aixinge/utils"
	"aixinge/utils/snowflake"
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

func (c *ApplicationService) Delete(idsReq request.IdsReq) (err error) {
	err = global.DB.Delete(&[]message.Channel{}, "id in ?", idsReq.Ids).Error
	return err
}

func (c *ApplicationService) Update(app message.Application) (error, message.Application) {
	err := global.DB.Updates(&app).Error
	return err, app
}

func (c *ApplicationService) GetById(id snowflake.ID) (err error, mt message.Application) {
	err = global.DB.Where("id = ?", id).First(&mt).Error
	return err, mt
}

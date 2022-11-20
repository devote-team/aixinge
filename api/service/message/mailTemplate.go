package message

import (
	"aixinge/api/model/common/request"
	"aixinge/api/model/message"
	"aixinge/global"
	"aixinge/utils"
	"aixinge/utils/snowflake"
)

type MailTemplateService struct {
}

func (c *MailTemplateService) Create(app message.MailTemplate) (err error) {
	app.ID = utils.Id()
	// 状态，1、正常 2、禁用
	app.Status = 1
	err = global.DB.Create(&app).Error
	return err
}

func (c *MailTemplateService) Delete(idsReq request.IdsReq) (err error) {
	err = global.DB.Delete(&[]message.MailTemplate{}, "id in ?", idsReq.Ids).Error
	return err
}

func (c *MailTemplateService) Update(reqMt message.MailTemplate) (err error, mt message.MailTemplate) {
	err = global.DB.Updates(&reqMt).Error
	return err, reqMt
}

func (c *MailTemplateService) GetById(id snowflake.ID) (err error, mt message.MailTemplate) {
	err = global.DB.Where("id = ?", id).First(&mt).Error
	return err, mt
}

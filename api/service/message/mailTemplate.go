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

func (c *MailTemplateService) Create(app message.MailTemplate) error {
	app.ID = utils.Id()
	// 状态，1、正常 2、禁用
	app.Status = 1
	return global.DB.Create(&app).Error
}

func (c *MailTemplateService) Delete(idsReq request.IdsReq) error {
	return global.DB.Delete(&[]message.MailTemplate{}, "id in ?", idsReq.Ids).Error
}

func (c *MailTemplateService) Update(mt message.MailTemplate) (error, message.MailTemplate) {
	return global.DB.Updates(&mt).Error, mt
}

func (c *MailTemplateService) GetById(id snowflake.ID) (err error, mt message.MailTemplate) {
	err = global.DB.Where("id = ?", id).First(&mt).Error
	return err, mt
}

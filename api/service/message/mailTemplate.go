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

func (e *MailTemplateService) Create(app message.MailTemplate) error {
	app.ID = utils.Id()
	// 状态，1、正常 2、禁用
	app.Status = 1
	return global.DB.Create(&app).Error
}

func (e *MailTemplateService) Delete(idsReq request.IdsReq) error {
	return global.DB.Delete(&[]message.MailTemplate{}, "id in ?", idsReq.Ids).Error
}

func (e *MailTemplateService) Update(mt message.MailTemplate) (error, message.MailTemplate) {
	return global.DB.Updates(&mt).Error, mt
}

func (e *MailTemplateService) GetById(id snowflake.ID) (err error, mt message.MailTemplate) {
	err = global.DB.Where("id = ?", id).First(&mt).Error
	return err, mt
}

func (e *MailTemplateService) Page(page request.PageInfo) (err error, list interface{}, total int64) {
	db := global.DB.Model(&message.MailTemplate{})
	var mtList []message.MailTemplate
	err = db.Count(&total).Error
	if total > 0 {
		err = db.Limit(page.PageSize).Offset(page.Offset()).Find(&mtList).Error
	}
	return err, mtList, total
}

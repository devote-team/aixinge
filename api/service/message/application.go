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

func (a *ApplicationService) Create(app message.Application) error {
	app.ID = utils.Id()
	// 状态，1、正常 2、禁用
	app.Status = 1
	return global.DB.Create(&app).Error
}

func (a *ApplicationService) Delete(idsReq request.IdsReq) error {
	return global.DB.Delete(&[]message.Channel{}, "id in ?", idsReq.Ids).Error
}

func (a *ApplicationService) Update(app message.Application) (error, message.Application) {
	return global.DB.Updates(&app).Error, app
}

func (a *ApplicationService) GetById(id snowflake.ID) (err error, mt message.Application) {
	err = global.DB.Where("id = ?", id).First(&mt).Error
	return err, mt
}

func (a *ApplicationService) Page(page request.PageInfo) (err error, list interface{}, total int64) {
	db := global.DB.Model(&message.Application{})
	var appList []message.Application
	err = db.Count(&total).Error
	if total > 0 {
		err = db.Limit(page.PageSize).Offset(page.Offset()).Find(&appList).Error
	}
	return err, appList, total
}

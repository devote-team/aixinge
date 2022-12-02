package message

import (
	"aixinge/api/model/common/request"
	"aixinge/api/model/message"
	"aixinge/global"
	"aixinge/utils"
)

type MailLogService struct {
}

func (m *MailLogService) Create(ml message.MailLog) error {
	ml.ID = utils.Id()
	return global.DB.Create(&ml).Error
}

func (m *MailLogService) Delete(idsReq request.IdsReq) error {
	return global.DB.Delete(&[]message.MailLog{}, "id in ?", idsReq.Ids).Error
}

func (m *MailLogService) Page(page request.PageInfo) (err error, list interface{}, total int64) {
	db := global.DB.Model(&message.MailLog{})
	var mlList []message.MailLog
	err = db.Count(&total).Error
	if total > 0 {
		err = db.Limit(page.PageSize).Offset(page.Offset()).Find(&mlList).Error
	}
	return err, mlList, total
}

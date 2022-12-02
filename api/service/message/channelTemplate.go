package message

import (
	"aixinge/api/model/common/request"
	"aixinge/api/model/message"
	"aixinge/global"
)

type ChannelTemplateService struct {
}

func (c *ChannelTemplateService) Create(ct message.ChannelTemplate) error {
	return global.DB.Create(&ct).Error
}

func (c *ChannelTemplateService) Delete(idsReq request.IdsReq) error {
	return global.DB.Delete(&[]message.Channel{}, "id in ?", idsReq.Ids).Error
}

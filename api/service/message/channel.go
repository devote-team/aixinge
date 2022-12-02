package message

import (
	"aixinge/api/model/common/request"
	"aixinge/api/model/message"
	"aixinge/global"
	"aixinge/utils"
	"aixinge/utils/snowflake"
)

type ChannelService struct {
}

func (c *ChannelService) Create(channel message.Channel) error {
	channel.ID = utils.Id()
	// 状态，1、正常 2、禁用
	channel.Status = 1
	return global.DB.Create(&channel).Error
}

func (c *ChannelService) Delete(idsReq request.IdsReq) error {
	return global.DB.Delete(&[]message.Channel{}, "id in ?", idsReq.Ids).Error
}

func (c *ChannelService) Update(channel message.Channel) (error, message.Channel) {
	return global.DB.Updates(&channel).Error, channel
}

func (c *ChannelService) GetById(id snowflake.ID) (err error, mt message.Channel) {
	err = global.DB.Where("id = ?", id).First(&mt).Error
	return err, mt
}

func (c *ChannelService) Page(page request.PageInfo) (err error, list interface{}, total int64) {
	db := global.DB.Model(&message.Channel{})
	var channelList []message.Channel
	err = db.Count(&total).Error
	if total > 0 {
		err = db.Limit(page.PageSize).Offset(page.Offset()).Find(&channelList).Error
	}
	return err, channelList, total
}

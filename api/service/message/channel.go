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

func (c *ChannelService) Create(channel message.Channel) (err error) {
	channel.ID = utils.Id()
	// 状态，1、正常 2、禁用
	channel.Status = 1
	err = global.DB.Create(&channel).Error
	return err
}

func (c *ChannelService) Delete(idsReq request.IdsReq) (err error) {
	err = global.DB.Delete(&[]message.Channel{}, "id in ?", idsReq.Ids).Error
	return err
}

func (c *ChannelService) Update(reqChannel message.Channel) (error, message.Channel) {
	err := global.DB.Updates(&reqChannel).Error
	return err, reqChannel
}

func (c *ChannelService) GetById(id snowflake.ID) (err error, mt message.Channel) {
	err = global.DB.Where("id = ?", id).First(&mt).Error
	return err, mt
}

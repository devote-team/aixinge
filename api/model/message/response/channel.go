package response

import (
	"aixinge/api/model/message"
)

type ChannelResponse struct {
	Channel message.Channel `json:"channel"`
}

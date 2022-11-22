package response

import "aixinge/api/model/message"

type AppResponse struct {
	Application message.Application `json:"application"`
}

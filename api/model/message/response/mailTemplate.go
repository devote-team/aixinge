package response

import (
	"aixinge/api/model/message"
)

type MailTemplateResponse struct {
	MailTemplate message.MailTemplate `json:"mailTemplate"`
}

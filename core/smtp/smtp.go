package smtp

import (
	"github.com/emersion/go-sasl"
	"github.com/emersion/go-smtp"
	"strings"
)

type Smtp struct {
	address string
	client  sasl.Client
}

func NewSmtp(identity, address, username, password string) Smtp {
	s := Smtp{}
	s.address = address
	s.client = sasl.NewPlainClient(identity, username, password)
	return s
}

func (s Smtp) SendMail(from string, to []string, subject, content string) error {
	return smtp.SendMail(s.address, s.client, from, to, strings.NewReader("To: "+strings.Join(to, ",")+"\r\n"+
		"Subject: "+subject+"\r\n"+
		"\r\n"+content+"\r\n"))
}

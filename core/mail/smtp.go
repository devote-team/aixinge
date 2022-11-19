package mail

import (
	"bytes"
	"net/smtp"
	"strings"
)

type Smtp struct {
	address  string
	username string
	auth     smtp.Auth
}

func NewSmtp(identity, address, username, password string) Smtp {
	s := Smtp{}
	hp := strings.Split(address, ":")
	s.username = username
	s.auth = smtp.PlainAuth(identity, username, password, hp[0])
	return s
}

// SendMail 发送邮件
// html 发送内容是否为 Html
// subject 主题
// content 内容
// replyToAddress 回信地址
// to 收件人
// cc 抄送人
// bcc 密送人
func (s Smtp) SendMail(html bool, subject, content, replyToAddress string, to, cc, bcc []string) error {
	var sendTo = to
	var bt bytes.Buffer
	// 收件人
	bt.WriteString("To:")
	bt.WriteString(strings.Join(to, ","))
	bt.WriteString("\r\n")
	// 主题
	bt.WriteString("Subject: ")
	bt.WriteString(subject)
	bt.WriteString("\r\n")
	// 回信地址
	bt.WriteString("Reply-To: ")
	bt.WriteString(replyToAddress)
	bt.WriteString("\r\n")
	// 抄送人
	if len(cc) > 0 {
		sendTo = append(sendTo, cc...)
		bt.WriteString("Cc: ")
		bt.WriteString(strings.Join(cc, ","))
		bt.WriteString("\r\n")
	}
	// 密送人
	if len(bcc) > 0 {
		sendTo = append(sendTo, bcc...)
		bt.WriteString("Bcc: ")
		bt.WriteString(strings.Join(bcc, ","))
		bt.WriteString("\r\n")
	}
	// 内容类型
	bt.WriteString("Content-Type: text/")
	if html {
		bt.WriteString("html")
	} else {
		bt.WriteString("plain")
	}
	bt.WriteString("; charset=UTF-8")
	bt.WriteString("\r\n\r\n")
	bt.WriteString(content)
	return smtp.SendMail(s.address, s.auth, s.username, sendTo, bt.Bytes())
}

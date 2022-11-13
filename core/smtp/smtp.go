package smtp

import (
	"fmt"
	"net/smtp"
	"strings"
)

func SendMail(user, password, host, subject, content, mailType, replyToAddress string, to, cc, bcc []string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	var contentType string
	if mailType == "html" {
		contentType = "Content-Type: text/" + mailType + "; charset=UTF-8"
	} else {
		contentType = "Content-Type: text/plain" + "; charset=UTF-8"
	}
	toAddress := strings.Join(to, ",")
	ccAddress := strings.Join(cc, ",")
	bccAddress := strings.Join(bcc, ",")
	msg := []byte("To: " + toAddress + "\r\n" +
		"From: " + user + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"Reply-To: " + replyToAddress + "\r\n" +
		"Cc: " + ccAddress + "\r\n" +
		"Bcc: " + bccAddress + "\r\n" +
		contentType + "\r\n\r\n" +
		content)
	sendTo := Merge(to, cc, bcc)
	err := smtp.SendMail(host, auth, user, sendTo, msg)
	if err != nil {
		fmt.Printf("send smtp mail error: %v", err)
	}
	return err
}

func Merge(to []string, cc []string, bcc []string) []string {
	return append(append(to, cc...), bcc...)
}

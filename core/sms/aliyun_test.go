package sms

import (
	"fmt"
	"testing"
)

func TestSendSms(t *testing.T) {
	accessKeyId := "accessKeyId"
	accessKeySecret := "accessKeySecret"
	phoneNumbers := []string{
		"电话号码",
	}
	signName := "短信签名"
	templateCode := "短信模板"
	templateParam := "{\"code\":\"6683\"}" // 短信参数

	client := CreateClient(accessKeyId, accessKeySecret)
	response := client.SendSms(phoneNumbers, signName, templateCode, templateParam)
	fmt.Printf("%+v\n", response)
}

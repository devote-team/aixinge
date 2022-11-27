package sms

import (
	"aixinge/global"
	"aixinge/utils/helper"
	"encoding/json"
	"go.uber.org/zap"
	"io"
	"net/http"
	"strings"
)

const (
	Endpoint = "dysmsapi.aliyuncs.com"
	GET      = "GET"
	POST     = "POST"
	Version  = "2017-05-25"
)

type AliyunSmsClient struct {
	accessKeyId     string
	accessKeySecret string
}

type AliyunSmsResponse struct {
	Message   string `json:"message"`
	RequestId string `json:"requestId"`
	Code      string `json:"code"`
	BizId     string `json:"bizId"`
}

func CreateClient(accessKeyId, accessKeySecret string) AliyunSmsClient {
	c := AliyunSmsClient{}
	c.accessKeyId = accessKeyId
	c.accessKeySecret = accessKeySecret
	return c
}

func (c AliyunSmsClient) SendSms(phoneNumbers []string, signName, templateCode, templateParam string) AliyunSmsResponse {
	parameters := map[string]string{
		"PhoneNumbers":  strings.Join(phoneNumbers, ","),
		"SignName":      signName,
		"TemplateCode":  templateCode,
		"TemplateParam": templateParam,
	}
	url := helper.BuildOpenApiRequestUrl("SendSms", Version, GET, Endpoint, c.accessKeyId, c.accessKeySecret, parameters)
	resp, requestErr := http.Get(url)
	if requestErr != nil {
		global.LOG.Error("请求发送阿里云短信失败！", zap.Any("err", requestErr))
	}
	defer resp.Body.Close()
	body, readErr := io.ReadAll(resp.Body)
	if readErr != nil {
		global.LOG.Error("解析阿里云短信响应失败！", zap.Any("err", readErr))
	}

	smsResponse := AliyunSmsResponse{}
	jsonErr := json.Unmarshal(body, &smsResponse)
	if jsonErr != nil {
		global.LOG.Error("解析阿里云短信响应失败！", zap.Any("err", jsonErr))
	}
	return smsResponse
}

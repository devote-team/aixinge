package helper

import (
	"aixinge/utils"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"net/url"
	"sort"
	"strings"
	"time"
)

func GetUtcTime() string {
	now := time.Now()
	year, mon, day := now.UTC().Date()
	hour, min, sec := now.UTC().Clock()
	s := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02dZ", year, mon, day, hour, min, sec)
	return s
}

func PercentEncode(str string) string {
	str = strings.Replace(str, "+", "%20", -1)
	str = strings.Replace(str, "*", "%2A", -1)
	str = strings.Replace(str, "%7E", "~", -1)
	return str
}

func InitCommonRequestParameters(accessKeyId, action, version string) map[string]string {
	var commonParameters = map[string]string{
		"Format":           "JSON",
		"SignatureMethod":  "HMAC-SHA1",
		"SignatureVersion": "1.0",
	}
	commonParameters["Action"] = action
	commonParameters["AccessKeyId"] = accessKeyId
	commonParameters["Version"] = version
	commonParameters["SignatureNonce"] = utils.Id().String()
	commonParameters["Timestamp"] = GetUtcTime()
	return commonParameters
}

func BuildUrlParams(params map[string]string) url.Values {
	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	urlParams := url.Values{}
	for _, k := range keys {
		urlParams.Add(k, params[k])
	}

	return urlParams
}

func BuildSignStr(requestType, standardRequestStr string) string {
	return requestType + "&" + url.QueryEscape("/") + "&" + "" + url.QueryEscape(standardRequestStr)
}

func BuildSignature(accessKeySecret, signStr string) string {
	key := []byte(accessKeySecret + "&")
	mac := hmac.New(sha1.New, key)
	mac.Write([]byte(signStr))
	res := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	return res
}

func BuildOpenApiRequestUrl(action, version, requestType, endpoint, accessKeyId, accessKeySecret string, params map[string]string) string {
	requestParameters := InitCommonRequestParameters(accessKeyId, action, version)
	for k, v := range params {
		requestParameters[k] = v
	}

	urlParams := BuildUrlParams(requestParameters)
	encodeUrlParams := urlParams.Encode()
	percent := PercentEncode(encodeUrlParams)
	signStr := BuildSignStr(requestType, percent)
	signature := BuildSignature(accessKeySecret, signStr)
	return "https://" + endpoint + "/?" + encodeUrlParams + "&Signature=" + signature
}

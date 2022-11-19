package helper

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestAliyunApiUtil(t *testing.T) {
	parameters := InitCommonRequestParameters("123", "SendSms", "2017-05-25")
	out, _ := json.Marshal(parameters)
	fmt.Println("Parameters：")
	fmt.Println(string(out))
	fmt.Println()

	urlParams := BuildUrlParams(parameters)
	fmt.Println("StandardRequestStr：")
	fmt.Println(urlParams)
	fmt.Println()

	encodeUrlParams := urlParams.Encode()
	percent := PercentEncode(encodeUrlParams)
	signStr := BuildSignStr("GET", percent)
	fmt.Println("SignStr：")
	fmt.Println(signStr)
	fmt.Println()

	signature := BuildSignature("test", signStr)
	fmt.Println("Signature：")
	fmt.Println(signature)
	fmt.Println()
}

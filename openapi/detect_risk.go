package openapi

import (
	"net/http"

	"github.com/sanxibengbeng/swan-sdk-go/utils"
)

type DetectRiskParam struct {
	AccessToken string
	Appkey      string
	Xtoken      string
	Clientip    string
	Ev          string
	Useragent   string
	Phone       string
	Ts          string
	Type        string
}

type DetectRiskResp struct {
	Errno     int    `json:"errno"`
	ErrorMsg  string `json:"error_msg"`
	ErrorCode int    `json:"error_code"`
	Errmsg    string `json:"errmsg"`
	RequestId string `json:"request_id"`
	Timestamp int64  `json:"timestamp"`
	Data      struct {
		Level string   `json:"level"`
		Tag   []string `json:"tag"`
	} `json:"data"`
}

// DectectRisk 发起检测用户是否是作弊用户请求
func DectectRisk(params *DetectRiskParam) (*DetectRiskResp, error) {
	var err error
	respData := &DetectRiskResp{}

	client := utils.NewHTTPClient().
		SetContentType(utils.ContentTypeForm).
		SetConverterType(utils.ConverterTypeJSON).
		SetMethod(http.MethodPost).
		SetScheme("https").
		SetHost("openapi.baidu.com").
		SetPath("/rest/2.0/smartapp/detectrisk")
	client.AddGetParam("access_token", params.AccessToken)

	client.AddPostParam("appkey", params.Appkey)
	client.AddPostParam("xtoken", params.Xtoken)
	client.AddPostParam("type", params.Type)
	client.AddPostParam("clientip", params.Clientip)
	client.AddPostParam("ts", params.Ts)
	client.AddPostParam("ev", params.Ev)
	client.AddPostParam("useragent", params.Useragent)
	client.AddPostParam("phone", params.Phone)

	err = client.Do()
	if err != nil {
		return nil, err
	}
	err = client.Convert(respData)
	return respData, err
}

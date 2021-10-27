package swansdk

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

	client := newHTTPClient().
		setContentType(contentTypeForm).
		setConverterType(converterTypeJSON).
		setMethod(methodPOST).
		setScheme("https").
		setHost("openapi.baidu.com").
		setPath("/rest/2.0/smartapp/detectrisk")

	client.addGetParam("access_token", params.AccessToken)

	client.addPostParam("appkey", params.Appkey)
	client.addPostParam("xtoken", params.Xtoken)
	client.addPostParam("type", params.Type)
	client.addPostParam("clientip", params.Clientip)
	client.addPostParam("ts", params.Ts)
	client.addPostParam("ev", params.Ev)
	client.addPostParam("useragent", params.Useragent)
	client.addPostParam("phone", params.Phone)

	err = client.do()
	if err != nil {
		return nil, err
	}
	err = client.convert(respData)
	return respData, err
}

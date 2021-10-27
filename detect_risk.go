package swansdk

type detectRiskParam struct {
	accessToken string
	appkey      string
	xtoken      string
	clientip    string
	ev          string
	useragent   string
	phone       string
	ts          string
	_type       string
}

type detectRiskResp struct {
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

// dectectRisk 发起检测用户是否是作弊用户请求
func dectectRisk(params *detectRiskParam) (*detectRiskResp, error) {
	var err error
	respData := &detectRiskResp{}

	client := newHTTPClient().
		setContentType(contentTypeForm).
		setScheme("https").
		setHost("openapi.baidu.com").
		setPath("/rest/2.0/smartapp/detectrisk")

	client.addGetParam("access_token", params.accessToken)

	client.addPostParam("appkey", params.appkey)
	client.addPostParam("xtoken", params.xtoken)
	client.addPostParam("type", params._type)
	client.addPostParam("clientip", params.clientip)
	client.addPostParam("ts", params.ts)
	client.addPostParam("ev", params.ev)
	client.addPostParam("useragent", params.useragent)
	client.addPostParam("phone", params.phone)

	err = client.do()
	if err != nil {
		return nil, err
	}
	err = client.convert(respData)
	return respData, err
}

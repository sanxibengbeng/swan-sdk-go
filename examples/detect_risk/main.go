package main

import (
	"fmt"

	"github.com/sanxibengbeng/swan-sdk-go/swansdk"
)

func main() {
	// 获取方式参考文档中 "post 参数" 解释
	//xtokenMap := map[string]string{
	//	"key":   "XXX",
	//	"value": "XXX",
	//}
	//xtokenByte, err := json.Marshal(xtokenMap)
	//if err != nil {
	//	log.Fatalln(err)
	//	return
	//}
	//// 参考 accessToken 获取文档；
	accessToken := "xxx"
	//// 小程序 appKey,从 B 端平台获取；
	//appKey := "xxxx"
	//// 用户 ip
	//clientIp := "xxx.xxx.xxx.xxx"
	//// 11位明文手机号
	//phone := "130xxxxxxxx"
	//// 11位明文手机号
	//useragent := "xxxxxxx"
	//// ev 1 为点击按钮获取，请按照文档设置合理的数值
	//ev := "1"
	//xtoken := string(xtokenByte)
	swansdk.Debug()
	reqParams := &swansdk.DetectRiskParam{
		AccessToken: accessToken,
		Useragent:   "xxxx",
	}
	resp, err := swansdk.DectectRisk(reqParams)
	fmt.Printf("%#v[%#v]", resp, err)
}

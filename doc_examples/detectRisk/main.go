// 本示例基于百度智能小程序服务端开发者 OpenAPI-SDK-Go
// 使用该示例需要首先下载该 SDK，使用引导见：https://smartprogram.baidu.com/docs/develop/serverapi/introduction_for_openapi_sdk/
// 使用之前请先确认下 SDK 版本是否为最新版本，如不是，请下载最新版本使用
// 如使用过程中遇到问题，可以加入如流群：5702992，进行反馈咨询
package main

import (
	"fmt"

	// github.com/baidu/smartapp-openapi-go/openapi 为示例项目名，实际使用时需要替换为实际的项目名
	"github.com/baidu/smartapp-openapi-go/openapi"
)

func main() {
	// 开发者在此设置请求参数，文档示例中的参数均为示例参数，实际参数请参考对应接口的文档上方的参数说明填写
	// 注意：代码示例中的参数字段基本是驼峰形式，而文档中的参数说明的参数字段基本是下划线形式
	// 如果开发者不想传非必需参数，可以将设置该参数的行注释
	reqParams := &openapi.DetectRiskRequest{
		AccessToken: "20.851f83125a4622a8c7ad1c144020eb57.7884168.3723216600.467842-11427304",                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  // 文档中对应字段：access_token，实际使用时请替换成真实参数
		Appkey:      "WXF1pGOvo1TTGU4qCMMhEjvFBkF0bO0Z",                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        // 文档中对应字段：appkey，实际使用时请替换成真实参数
		Xtoken:      "{\"key\":\"jU+mx3VJ6+k8+JfN8cLPNfQZIVbCAZwhMIlTtnt8yl+YTPK7E+58s48UwTqR2eKQEyKu0Qbd4xknxNNoYl7w7o/6/qemfUNn1pDXmsYnaZz8tM6k1bhRD7TusfOXXqXRo8gWuUdnWttnIhxvYKGwhzL5sUF3fqnxY5S2PUnGE2g=\", \"value\": \"TPHtjm7RTDX0pUpcUjbhRu/t4MA60kF+mFv7DPmNSx4zMsTsT7Yitu+DoQ6CJS3f4tQBHpqzQ1vfW7nV8Zm7HWkkXK2xkF3jSTSEWH7KkLAMdzWwqLKZQTaWG0r+MU+1qOqYF7mc22oB4WSSfPJQ5ZUYpY+7RezUMWK4xyUB/4vEy53HZ4SYZjsfmJOYNcVsh3A6fTsoHDsNBiXYA3KUe8ZxiSzmyLYe7EYjW2XLcL+iUgcToNuH468Ypn+Py7OxOD8lS2BgWVNV6sdGriYuRDAN8rcugPbVscFoEeOcDWIDaHNKs751vDvmQQCc5M7EXsQ3W/NDdze50dgJ5AL8ZLV+5Ahe1ISoxflpRKjvl6Jl10+p8jESon6DLJA86/+n3FAbCifa8mZLvyHJ+gTSR0h3lLSZW3ZntrbeofVP4MZTYsPip3k1Kt4A6G/ABj7K3k6FIx1iM4UQWvPgFFOJ/vbCf0c6FXVDLHDid8V5qGwJ4TTRur2MJH0yVPiS0dltOQkIIAQcK7C+nTgi+EKY8RwwoOYw\"}", // 文档中对应字段：xtoken，实际使用时请替换成真实参数
		Type:        "marketing",                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               // 文档中对应字段：type，实际使用时请替换成真实参数
		Clientip:    "128.1.8.7",                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               // 文档中对应字段：clientip，实际使用时请替换成真实参数
		Ts:          1445465721,                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                // 文档中对应字段：ts，实际使用时请替换成真实参数
		Ev:          "1",                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       // 文档中对应字段：ev，实际使用时请替换成真实参数
		Useragent:   "Mozilla/1.4 (Macintosh, Intel Mac OS X 54_30_1) AppleWebKit/515.68 (KHTML, like Gecko) Chrome/86.3.7115.87 Safari/036.25",                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                // 文档中对应字段：useragent，实际使用时请替换成真实参数
		Phone:       "",                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        // 文档中对应字段：phone，实际使用时请替换成真实参数
	}

	resp, err := openapi.DetectRisk(reqParams)
	if err != nil {
		if _, ok := err.(*openapi.OpenAPIError); ok {
			// openapi error
			// 可能是 access_token 无效，可以尝试重新生成 access_token
			fmt.Println("openapi error, ", err)
		} else if _, ok := err.(*openapi.APIError); ok {
			// api error
			// 可能是参数错误或没有权限，建议根据错误信息自查或者社区发帖、加群反馈
			fmt.Println("api error, ", err)
		} else {
			// 其他错误
			fmt.Println("others error, ", err)
		}
	} else {
		fmt.Printf("%#v\n", *resp)
	}
}

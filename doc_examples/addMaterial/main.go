// 本示例基于百度智能小程序服务端开发者 OpenAPI-SDK-Go
// 使用该示例需要首先下载该 SDK，使用引导见：https://smartprogram.baidu.com/docs/develop/serverapi/introduction_for_openapi_sdk/
// 使用之前请先确认下 SDK 版本是否为最新版本，如不是，请下载最新版本使用
// 如使用过程中遇到问题，可以加入如流群：5702992，进行反馈咨询
package main

import (
	"fmt"

	"github.com/baidu/smartapp-openapi-go/openapi"
)

func main() {
	// 开发者在此设置请求参数，文档示例中的参数均为示例参数，实际参数请参考对应接口的文档上方的参数说明填写
	// 注意：代码示例中的参数字段基本是驼峰形式，而文档中的参数说明的参数字段基本是下划线形式
	// 如果开发者不想传非必需参数，可以将设置该参数的行注释
	reqParams := &openapi.AddMaterialRequest{
		AccessToken:   "28.0fa855ebfa3c181dda3b8a783c23520e.4567236.1351151770.513624-83563323",                                  // 文档中对应字段：access_token，实际使用时请替换成真实参数
		AppID:         17071416,                                                                                                  // 文档中对应字段：app_id，实际使用时请替换成真实参数
		ImageURL:      "https://mbs1.bdstatic.com/searchbox/mappconsole/image/45326405/d32d8118-abf5-4d27-a4c1-df1e2b04645b.png", // 文档中对应字段：imageUrl，实际使用时请替换成真实参数
		Title:         "测试数据",                                                                                                    // 文档中对应字段：title，实际使用时请替换成真实参数
		Path:          "/pages/index/index",                                                                                      // 文档中对应字段：path，实际使用时请替换成真实参数
		Category1Code: "7",                                                                                                       // 文档中对应字段：category1Code，实际使用时请替换成真实参数
		Category2Code: "70255",                                                                                                   // 文档中对应字段：category2Code，实际使用时请替换成真实参数
		Desc:          "",                                                                                                        // 文档中对应字段：desc，实际使用时请替换成真实参数
		LabelAttr:     "",                                                                                                        // 文档中对应字段：labelAttr，实际使用时请替换成真实参数
		LabelDiscount: "",                                                                                                        // 文档中对应字段：labelDiscount，实际使用时请替换成真实参数
		ButtonName:    "",                                                                                                        // 文档中对应字段：buttonName，实际使用时请替换成真实参数
		BigImage:      "",                                                                                                        // 文档中对应字段：bigImage，实际使用时请替换成真实参数
		VerticalImage: "",                                                                                                        // 文档中对应字段：verticalImage，实际使用时请替换成真实参数
		ExtJSON:       "",                                                                                                        // 文档中对应字段：extJson，实际使用时请替换成真实参数
	}

	resp, err := openapi.AddMaterial(reqParams)
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

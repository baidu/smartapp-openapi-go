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
	reqParams := &openapi.ModifyMaterialRequest{
		AccessToken:   "22.2fa733ebfa5c216dda3b8a428c20764e.4464173.0458024105.803580-47700651",                                  // 文档中对应字段：access_token，实际使用时请替换成真实参数
		AppID:         13375626,                                                                                                  // 文档中对应字段：app_id，实际使用时请替换成真实参数
		ID:            13654,                                                                                                     // 文档中对应字段：id，实际使用时请替换成真实参数
		ImageURL:      "https://mbs0.bdstatic.com/searchbox/mappconsole/image/35037458/d65d3103-abf2-5d35-a0c3-df8e8b00827b.png", // 文档中对应字段：imageUrl，实际使用时请替换成真实参数
		Title:         "修改测试数据",                                                                                                  // 文档中对应字段：title，实际使用时请替换成真实参数
		Path:          "/pages/index/index",                                                                                      // 文档中对应字段：path，实际使用时请替换成真实参数
		Category1Code: "7",                                                                                                       // 文档中对应字段：category1Code，实际使用时请替换成真实参数
		Category2Code: "72142",                                                                                                   // 文档中对应字段：category2Code，实际使用时请替换成真实参数
		Desc:          "",                                                                                                        // 文档中对应字段：desc，实际使用时请替换成真实参数
		LabelAttr:     "",                                                                                                        // 文档中对应字段：labelAttr，实际使用时请替换成真实参数
		LabelDiscount: "",                                                                                                        // 文档中对应字段：labelDiscount，实际使用时请替换成真实参数
		ButtonName:    "",                                                                                                        // 文档中对应字段：buttonName，实际使用时请替换成真实参数
		BigImage:      "",                                                                                                        // 文档中对应字段：bigImage，实际使用时请替换成真实参数
		VerticalImage: "",                                                                                                        // 文档中对应字段：verticalImage，实际使用时请替换成真实参数
		ExtJSON:       "",                                                                                                        // 文档中对应字段：extJson，实际使用时请替换成真实参数
	}

	resp, err := openapi.ModifyMaterial(reqParams)
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
		fmt.Printf("%#v\n", resp)
	}
}

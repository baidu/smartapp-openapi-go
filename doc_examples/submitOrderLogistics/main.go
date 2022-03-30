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
	SubmitOrderLogisticsRequestDataItemEXTMainOrderExpressVal := openapi.SubmitOrderLogisticsRequestDataItemEXTMainOrderExpress{
		Code:   "SFEXPRESS",  // 文档中对应字段：Code，实际使用时请替换成真实参数
		ID:     "1877828875", // 文档中对应字段：ID，实际使用时请替换成真实参数
		Name:   "顺丰快递",       // 文档中对应字段：Name，实际使用时请替换成真实参数
		Status: 0,            // 文档中对应字段：Status，实际使用时请替换成真实参数
		Type:   1,            // 文档中对应字段：Type，实际使用时请替换成真实参数
	}

	SubmitOrderLogisticsRequestDataItemEXTMainOrderVal := openapi.SubmitOrderLogisticsRequestDataItemEXTMainOrder{
		Express: SubmitOrderLogisticsRequestDataItemEXTMainOrderExpressVal, // 文档中对应字段：Express，实际使用时请替换成真实参数
	}

	SubmitOrderLogisticsRequestDataItemEXTVal := openapi.SubmitOrderLogisticsRequestDataItemEXT{
		MainOrder: SubmitOrderLogisticsRequestDataItemEXTMainOrderVal, // 文档中对应字段：MainOrder，实际使用时请替换成真实参数
	}

	SubmitOrderLogisticsRequestDataItemVal := openapi.SubmitOrderLogisticsRequestDataItem{
		BizAPPID:   "WXF8pGOvo7TTGU2qCMMhEjvFBkF5bO6Z",        // 文档中对应字段：BizAPPID，实际使用时请替换成真实参数
		CateID:     2,                                         // 文档中对应字段：CateID，实际使用时请替换成真实参数
		EXT:        SubmitOrderLogisticsRequestDataItemEXTVal, // 文档中对应字段：EXT，实际使用时请替换成真实参数
		ResourceID: "2135682501336",                           // 文档中对应字段：ResourceID，实际使用时请替换成真实参数
	}

	reqParams := &openapi.SubmitOrderLogisticsRequest{
		AccessToken: "21.655f01013a1620a0c6ad0c023114eb86.2737666.8284861865.634214-60374587", // 文档中对应字段：access_token，实际使用时请替换成真实参数
		OpenID:      "k64HEREQhWhWWB4WYqYT7ITUGX",                                             // 文档中对应字段：open_id，实际使用时请替换成真实参数
		SceneID:     "1255828404632",                                                          // 文档中对应字段：scene_id，实际使用时请替换成真实参数
		SceneType:   2,                                                                        // 文档中对应字段：scene_type，实际使用时请替换成真实参数
		PmAppKey:    "WXF0pGOvo8TTGU6qCMMhEjvFBkF3bO4Z",                                       // 文档中对应字段：pm_app_key，实际使用时请替换成真实参数
		Data: []openapi.SubmitOrderLogisticsRequestDataItem{
			SubmitOrderLogisticsRequestDataItemVal,
		}, // 文档中对应字段：Data，实际使用时请替换成真实参数
	}

	resp, err := openapi.SubmitOrderLogistics(reqParams)
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

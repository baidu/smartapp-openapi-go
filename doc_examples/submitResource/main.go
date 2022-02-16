// 本示例基于百度智能小程序服务端开发者 OpenAPI-SDK-Go
// 使用该示例需要首先下载该 SDK，使用引导见：https://smartprogram.baidu.com/docs/develop/serverapi/introduction_for_openapi_sdk/
// 使用之前请先确认下 SDK 版本是否为最新版本，如不是，请下载最新版本使用
// 如使用过程中遇到问题，可以加入如流群：5702992，进行反馈咨询
package main

import (
	"fmt"

	// openapisdk-go.demo/openapi 为示例项目名，实际使用时需要替换为实际的项目名
	"openapisdk-go.demo/openapi"
)

func main() {
	// 开发者在此设置请求参数，文档示例中的参数均为示例参数，实际参数请参考对应接口的文档上方的参数说明填写
	// 注意：代码示例中的参数字段基本是驼峰形式，而文档中的参数说明的参数字段基本是下划线形式
	// 如果开发者不想传非必需参数，可以将设置该参数的行注释
	reqParams := &openapi.SubmitResourceRequest{
		AccessToken: "25.7fa511ebfa6c515dda4b5a234c61272e.7631128.8088333066.483718-78502100", // 文档中对应字段：access_token，实际使用时请替换成真实参数
		AppID:       11751048,                                                                 // 文档中对应字段：app_id，实际使用时请替换成真实参数
		Body:        "黄油化开备用，黄油化开后加入糖霜，搅拌均匀，加入蛋清，继续打匀，加入切碎的蔓越莓，继续搅拌。蔓越莓放多少根据自己的喜...",          // 文档中对应字段：body，实际使用时请替换成真实参数
		Ext:         "{\"publish_time\": \"2021年 11 月 1 日\"}",                                 // 文档中对应字段：ext，实际使用时请替换成真实参数
		FeedSubType: "明星八卦（可选有限集合）",                                                           // 文档中对应字段：feed_sub_type，实际使用时请替换成真实参数
		FeedType:    "娱乐（可选有限集合）",                                                             // 文档中对应字段：feed_type，实际使用时请替换成真实参数
		Images:      "[\"https://z0.ax5x.com/6043/84/14/IP2kw4.jpg\"]",                        // 文档中对应字段：images，实际使用时请替换成真实参数
		MappSubType: "1558",                                                                   // 文档中对应字段：mapp_sub_type，实际使用时请替换成真实参数
		MappType:    "1071",                                                                   // 文档中对应字段：mapp_type，实际使用时请替换成真实参数
		Path:        "/pages/detail/detail?id=761446",                                         // 文档中对应字段：path，实际使用时请替换成真实参数
		Tags:        "电影",                                                                     // 文档中对应字段：tags，实际使用时请替换成真实参数
		Title:       "百度智能小程序，给你全新的智能体验",                                                      // 文档中对应字段：title，实际使用时请替换成真实参数
	}

	resp, err := openapi.SubmitResource(reqParams)
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

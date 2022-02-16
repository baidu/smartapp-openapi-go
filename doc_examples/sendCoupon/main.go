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
	reqParams := &openapi.SendCouponRequest{
		AccessToken:           "27.418f50154a0346a0c7ad2c838164eb21.3775265.3670804322.776682-23356153", // 文档中对应字段：access_token，实际使用时请替换成真实参数
		CouponID:              "8584413588",                                                             // 文档中对应字段：couponId，实际使用时请替换成真实参数
		OpenID:                "k37HEREQhWhWWB7WYqYT4ITUGX",                                             // 文档中对应字段：openId，实际使用时请替换成真实参数
		UniqueSendingIdentity: "xxxxx",                                                                  // 文档中对应字段：uniqueSendingIdentity，实际使用时请替换成真实参数
	}

	resp, err := openapi.SendCoupon(reqParams)
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

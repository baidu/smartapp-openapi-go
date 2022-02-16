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
	CreateCouponRequestbaseInfodateInfoVal := openapi.CreateCouponRequestbaseInfodateInfo{
		BeginTimestamp:    1277760566, // 文档中对应字段：beginTimestamp，实际使用时请替换成真实参数
		EndTimestamp:      1874286025, // 文档中对应字段：endTimestamp，实际使用时请替换成真实参数
		GetEndTimestamp:   1173315584, // 文档中对应字段：getEndTimestamp，实际使用时请替换成真实参数
		GetStartTimestamp: 1170541327, // 文档中对应字段：getStartTimestamp，实际使用时请替换成真实参数
		TimeUnit:          0,          // 文档中对应字段：timeUnit，实际使用时请替换成真实参数
		TimeValue:         0,          // 文档中对应字段：timeValue，实际使用时请替换成真实参数
		Type:              1,          // 文档中对应字段：type，实际使用时请替换成真实参数
	}

	CreateCouponRequestbaseInfoVal := openapi.CreateCouponRequestbaseInfo{
		AppRedirectPath: "/pages/index/index",                   // 文档中对应字段：appRedirectPath，实际使用时请替换成真实参数
		CodeType:        1,                                      // 文档中对应字段：codeType，实际使用时请替换成真实参数
		Color:           "B626",                                 // 文档中对应字段：color，实际使用时请替换成真实参数
		DateInfo:        CreateCouponRequestbaseInfodateInfoVal, // 文档中对应字段：dateInfo，实际使用时请替换成真实参数
		GetLimit:        3,                                      // 文档中对应字段：getLimit，实际使用时请替换成真实参数
		Quantity:        142,                                    // 文档中对应字段：quantity，实际使用时请替换成真实参数
		Title:           "自动化创建代金券858-63",                       // 文档中对应字段：title，实际使用时请替换成真实参数
	}

	reqParams := &openapi.CreateCouponRequest{
		AccessToken: "24.270f87852a2805a7c4ad5c264452eb16.4768545.0231116104.715673-85028007", // 文档中对应字段：access_token，实际使用时请替换成真实参数
		BaseInfo:    CreateCouponRequestbaseInfoVal,                                           // 文档中对应字段：baseInfo，实际使用时请替换成真实参数
		Description: "使用描述",                                                                   // 文档中对应字段：description，实际使用时请替换成真实参数
		CallbackURL: "/test",                                                                  // 文档中对应字段：callbackUrl，实际使用时请替换成真实参数
		CouponType:  "CASH",                                                                   // 文档中对应字段：couponType，实际使用时请替换成真实参数
		LeastCost:   12731,                                                                    // 文档中对应字段：leastCost，实际使用时请替换成真实参数
		ReduceCost:  1226,                                                                     // 文档中对应字段：reduceCost，实际使用时请替换成真实参数
		Discount:    "<no value>",                                                             // 文档中对应字段：discount，实际使用时请替换成真实参数
	}

	resp, err := openapi.CreateCoupon(reqParams)
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

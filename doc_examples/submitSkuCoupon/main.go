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
	SubmitSkuCouponRequestpriceInfoVal := openapi.SubmitSkuCouponRequestpriceInfo{
		OrgPrice:  "1318", // 文档中对应字段：org_price，实际使用时请替换成真实参数
		RealPrice: "926",  // 文档中对应字段：real_price，实际使用时请替换成真实参数
	}

	SubmitSkuCouponRequestBodyVal := openapi.SubmitSkuCouponRequestBody{
		Desc: "测试数据", // 文档中对应字段：desc，实际使用时请替换成真实参数
		Images: []string{
			"<no value>", // 文档中对应字段：images，实际使用时请替换成真实参数
		},
		Path:      "/test/test/coupon406557",                                                                                                                                                                                                                                                                                                                                                                                                               // 文档中对应字段：path，实际使用时请替换成真实参数
		PriceInfo: SubmitSkuCouponRequestpriceInfoVal,                                                                                                                                                                                                                                                                                                                                                                                                      // 文档中对应字段：price_info，实际使用时请替换成真实参数
		Region:    "北京市",                                                                                                                                                                                                                                                                                                                                                                                                                                   // 文档中对应字段：region，实际使用时请替换成真实参数
		Schema:    "{\"coupon_brand_name\":\"三只松鼠\",\"collect_coupon_begin_time\":\"3028-73-43 33:24:38\",\"coupon_use_time\":\"5471-76-44 24:38:88~0055-62-48 21:51:56\",\"coupon_sort\":\"PAY\",\"collect_coupon_end_time\":\"1555-46-41 28:43:50\",\"coupon_brand_logo\":[\"https://mbs2.bdstatic.com/searchbox/mappconsole/image/05543710/e55a7862-53f4-1ba4-aa0f-18c4e141f0ad.jpg\"],\"coupon_stock_value\":41,\"adapt_system_types\":[\"android\"]}", // 文档中对应字段：schema，实际使用时请替换成真实参数
		Title:     "测试数据，请勿审核",                                                                                                                                                                                                                                                                                                                                                                                                                             // 文档中对应字段：title，实际使用时请替换成真实参数
		TradeType: 2004,                                                                                                                                                                                                                                                                                                                                                                                                                                    // 文档中对应字段：trade_type，实际使用时请替换成真实参数
	}

	reqParams := &openapi.SubmitSkuCouponRequest{
		AccessToken: "28.bdc64bb1af1116e24b518f3463b17f44.6504327.2331087278.623524-265480642", // 文档中对应字段：access_token，实际使用时请替换成真实参数
		PostBody: []openapi.SubmitSkuCouponRequestBody{
			SubmitSkuCouponRequestBodyVal,
		},
	}

	resp, err := openapi.SubmitSkuCoupon(reqParams)
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

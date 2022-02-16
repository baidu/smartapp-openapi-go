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
	SubmitSkuRequestpriceInfoVal := openapi.SubmitSkuRequestpriceInfo{
		OrgPrice:      "334", // 文档中对应字段：org_price，实际使用时请替换成真实参数
		OrgUnit:       "天",   // 文档中对应字段：org_unit，实际使用时请替换成真实参数
		RangeMaxPrice: "",    // 文档中对应字段：range_max_price，实际使用时请替换成真实参数
		RangeMinPrice: "",    // 文档中对应字段：range_min_price，实际使用时请替换成真实参数
		RealPrice:     "274", // 文档中对应字段：real_price，实际使用时请替换成真实参数
	}

	SubmitSkuRequestactivityInfoItemVal := openapi.SubmitSkuRequestactivityInfoItem{
		ActivityDesc:      "满12减1",            // 文档中对应字段：activity_desc，实际使用时请替换成真实参数
		ActivityEndTime:   1608480000,         // 文档中对应字段：activity_end_time，实际使用时请替换成真实参数
		ActivityPath:      "/activity/coupon", // 文档中对应字段：activity_path，实际使用时请替换成真实参数
		ActivityStartTime: 1593985486,         // 文档中对应字段：activity_start_time，实际使用时请替换成真实参数
		ActivityType:      "领劵",               // 文档中对应字段：activity_type，实际使用时请替换成真实参数
	}

	SubmitSkuRequestBodyVal := openapi.SubmitSkuRequestBody{
		ActivityInfo: []openapi.SubmitSkuRequestactivityInfoItem{
			SubmitSkuRequestactivityInfoItemVal, // 文档中对应字段：activity_info，实际使用时请替换成真实参数
		},
		ButtonName: "预订/立即预订", // 文档中对应字段：button_name，实际使用时请替换成真实参数
		Desc:       "测试数据",    // 文档中对应字段：desc，实际使用时请替换成真实参数
		Images: []string{
			"<no value>", // 文档中对应字段：images，实际使用时请替换成真实参数
		},
		Path:      "/pages/detail/detail?id=025864",                                                                        // 文档中对应字段：path，实际使用时请替换成真实参数
		PriceInfo: SubmitSkuRequestpriceInfoVal,                                                                            // 文档中对应字段：price_info，实际使用时请替换成真实参数
		Region:    "北京市",                                                                                                   // 文档中对应字段：region，实际使用时请替换成真实参数
		Schema:    "{\"hotel_name\":\"北京大酒店\",\"hotel_addr\":\"北京西城区\",\"hotel_score\":\"2.3分\",\"hotel_star\":\"经济型酒店\"}", // 文档中对应字段：schema，实际使用时请替换成真实参数
		Tag:       "饮食健康;中餐",                                                                                               // 文档中对应字段：tag，实际使用时请替换成真实参数
		Title:     "测试数据，请勿审核",                                                                                             // 文档中对应字段：title，实际使用时请替换成真实参数
		TradeType: 5001,                                                                                                    // 文档中对应字段：trade_type，实际使用时请替换成真实参数
	}

	reqParams := &openapi.SubmitSkuRequest{
		AccessToken: "23.bdc37bb6af5564e67b184f2704b02f45.5532822.6302061577.883710-014501345", // 文档中对应字段：access_token，实际使用时请替换成真实参数
		PostBody: []openapi.SubmitSkuRequestBody{
			SubmitSkuRequestBodyVal,
		},
	}

	resp, err := openapi.SubmitSku(reqParams)
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

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
	reqParams := &openapi.ApplyOrderRefundRequest{
		AccessToken:      "23.200c0855c60c50787d3efd5f0863ebd5.2337260.5565488207.858408-14443121", // 文档中对应字段：access_token，实际使用时请替换成真实参数
		ApplyRefundMoney: 131,                                                                      // 文档中对应字段：applyRefundMoney，实际使用时请替换成真实参数
		BizRefundBatchID: "",                                                                       // 文档中对应字段：bizRefundBatchId，实际使用时请替换成真实参数
		IsSkipAudit:      0,                                                                        // 文档中对应字段：isSkipAudit，实际使用时请替换成真实参数
		OrderID:          1753737417222,                                                            // 文档中对应字段：orderId，实际使用时请替换成真实参数
		RefundReason:     "xxxxx",                                                                  // 文档中对应字段：refundReason，实际使用时请替换成真实参数
		RefundType:       1,                                                                        // 文档中对应字段：refundType，实际使用时请替换成真实参数
		TpOrderID:        "1764566176060",                                                          // 文档中对应字段：tpOrderId，实际使用时请替换成真实参数
		UserID:           1476884,                                                                  // 文档中对应字段：userId，实际使用时请替换成真实参数
		RefundNotifyURL:  "xxxxx",                                                                  // 文档中对应字段：refundNotifyUrl，实际使用时请替换成真实参数
		PmAppKey:         "MMUpGO",                                                                 // 文档中对应字段：pmAppKey，实际使用时请替换成真实参数
	}

	resp, err := openapi.ApplyOrderRefund(reqParams)
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

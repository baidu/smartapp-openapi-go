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
	AddOrderSubInfoRequestDataItemEXTSubsOrderItemsItemOrderDetailVal := openapi.AddOrderSubInfoRequestDataItemEXTSubsOrderItemsItemOrderDetail{
		Status:     "3",                                                                                                           // 文档中对应字段：Status，实际使用时请替换成真实参数
		SwanSchema: "baiduboxapp://swan/B1GF8AWvCSr17myIs74uqaoYz1pPCSY6/wjz/bdxd/order-detail/order-detail?orderId=727556018550", // 文档中对应字段：SwanSchema，实际使用时请替换成真实参数
	}

	AddOrderSubInfoRequestDataItemEXTSubsOrderItemsItemRefundProductItemVal := openapi.AddOrderSubInfoRequestDataItemEXTSubsOrderItemsItemRefundProductItem{
		Amount:   100,            // 文档中对应字段：Amount，实际使用时请替换成真实参数
		ID:       "152123021032", // 文档中对应字段：ID，实际使用时请替换成真实参数
		Quantity: 1,              // 文档中对应字段：Quantity，实际使用时请替换成真实参数
	}

	AddOrderSubInfoRequestDataItemEXTSubsOrderItemsItemRefundVal := openapi.AddOrderSubInfoRequestDataItemEXTSubsOrderItemsItemRefund{
		Amount: "1315", // 文档中对应字段：Amount，实际使用时请替换成真实参数
		Product: []openapi.AddOrderSubInfoRequestDataItemEXTSubsOrderItemsItemRefundProductItem{
			AddOrderSubInfoRequestDataItemEXTSubsOrderItemsItemRefundProductItemVal, // 文档中对应字段：Product，实际使用时请替换成真实参数
		},
	}

	AddOrderSubInfoRequestDataItemEXTSubsOrderItemsItemVal := openapi.AddOrderSubInfoRequestDataItemEXTSubsOrderItemsItem{
		CTime:       1571026203,                                                        // 文档中对应字段：CTime，实际使用时请替换成真实参数
		MTime:       1571026203,                                                        // 文档中对应字段：MTime，实际使用时请替换成真实参数
		OrderDetail: AddOrderSubInfoRequestDataItemEXTSubsOrderItemsItemOrderDetailVal, // 文档中对应字段：OrderDetail，实际使用时请替换成真实参数
		OrderType:   1,                                                                 // 文档中对应字段：OrderType，实际使用时请替换成真实参数
		Refund:      AddOrderSubInfoRequestDataItemEXTSubsOrderItemsItemRefundVal,      // 文档中对应字段：Refund，实际使用时请替换成真实参数
		SubOrderID:  "onlyOne",                                                         // 文档中对应字段：SubOrderID，实际使用时请替换成真实参数
		SubStatus:   401,                                                               // 文档中对应字段：SubStatus，实际使用时请替换成真实参数
	}

	AddOrderSubInfoRequestDataItemEXTSubsOrderVal := openapi.AddOrderSubInfoRequestDataItemEXTSubsOrder{
		Items: []openapi.AddOrderSubInfoRequestDataItemEXTSubsOrderItemsItem{
			AddOrderSubInfoRequestDataItemEXTSubsOrderItemsItemVal, // 文档中对应字段：Items，实际使用时请替换成真实参数
		},
		Status: 0, // 文档中对应字段：Status，实际使用时请替换成真实参数
	}

	AddOrderSubInfoRequestDataItemEXTVal := openapi.AddOrderSubInfoRequestDataItemEXT{
		SubsOrder: AddOrderSubInfoRequestDataItemEXTSubsOrderVal, // 文档中对应字段：SubsOrder，实际使用时请替换成真实参数
	}

	AddOrderSubInfoRequestDataItemVal := openapi.AddOrderSubInfoRequestDataItem{
		BizAPPID:   "WXF7pGOvo7TTGU1qCMMhEjvFBkF6bO6Z",   // 文档中对应字段：BizAPPID，实际使用时请替换成真实参数
		CateID:     1,                                    // 文档中对应字段：CateID，实际使用时请替换成真实参数
		EXT:        AddOrderSubInfoRequestDataItemEXTVal, // 文档中对应字段：EXT，实际使用时请替换成真实参数
		ResourceID: "1074262658687",                      // 文档中对应字段：ResourceID，实际使用时请替换成真实参数
	}

	reqParams := &openapi.AddOrderSubInfoRequest{
		AccessToken: "20.516f30737a5485a1c1ad5c630577eb17.6818331.1250101224.488872-48241650", // 文档中对应字段：access_token，实际使用时请替换成真实参数
		OpenID:      "k45HEREQhWhWWB5WYqYT2ITUGX",                                             // 文档中对应字段：open_id，实际使用时请替换成真实参数
		SceneID:     "1602430063447",                                                          // 文档中对应字段：scene_id，实际使用时请替换成真实参数
		SceneType:   2,                                                                        // 文档中对应字段：scene_type，实际使用时请替换成真实参数
		PmAppKey:    "WXF5pGOvo6TTGU4qCMMhEjvFBkF4bO4Z",                                       // 文档中对应字段：pm_app_key，实际使用时请替换成真实参数
		Data: []openapi.AddOrderSubInfoRequestDataItem{
			AddOrderSubInfoRequestDataItemVal,
		}, // 文档中对应字段：Data，实际使用时请替换成真实参数
	}

	resp, err := openapi.AddOrderSubInfo(reqParams)
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

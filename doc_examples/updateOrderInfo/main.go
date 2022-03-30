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
	UpdateOrderInfoRequestDataItemEXTMainOrderProductsItemSkuAttrItemVal := openapi.UpdateOrderInfoRequestDataItemEXTMainOrderProductsItemSkuAttrItem{
		Name:  "四川大凉山丑苹果脆甜",  // 文档中对应字段：Name，实际使用时请替换成真实参数
		Value: "5斤小果25个左右偏小", // 文档中对应字段：Value，实际使用时请替换成真实参数
	}

	UpdateOrderInfoRequestDataItemEXTMainOrderProductsItemDetailPageVal := openapi.UpdateOrderInfoRequestDataItemEXTMainOrderProductsItemDetailPage{
		Status:     2,                                                                                                             // 文档中对应字段：Status，实际使用时请替换成真实参数
		SwanSchema: "baiduboxapp://swan/B2GF1AWvCSr88myIs00uqaoYz4pPCSY2/wjz/bdxd/order-detail/order-detail?orderId=084242532606", // 文档中对应字段：SwanSchema，实际使用时请替换成真实参数
	}

	UpdateOrderInfoRequestDataItemEXTMainOrderProductsItemVal := openapi.UpdateOrderInfoRequestDataItemEXTMainOrderProductsItem{
		Desc:       "四川大凉山丑苹果脆甜:5斤小果25个左右偏小;",                                           // 文档中对应字段：Desc，实际使用时请替换成真实参数
		DetailPage: UpdateOrderInfoRequestDataItemEXTMainOrderProductsItemDetailPageVal, // 文档中对应字段：DetailPage，实际使用时请替换成真实参数
		ID:         "1438264586752",                                                     // 文档中对应字段：ID，实际使用时请替换成真实参数
		ImgList: []string{
			"<no value>", // 文档中对应字段：ImgList，实际使用时请替换成真实参数
		},
		Name:     "四川大凉山丑苹果脆甜红将军盐源丑苹果", // 文档中对应字段：Name，实际使用时请替换成真实参数
		PayPrice: 2390,                 // 文档中对应字段：PayPrice，实际使用时请替换成真实参数
		Price:    2390,                 // 文档中对应字段：Price，实际使用时请替换成真实参数
		Quantity: 1,                    // 文档中对应字段：Quantity，实际使用时请替换成真实参数
		SkuAttr: []openapi.UpdateOrderInfoRequestDataItemEXTMainOrderProductsItemSkuAttrItem{
			UpdateOrderInfoRequestDataItemEXTMainOrderProductsItemSkuAttrItemVal, // 文档中对应字段：SkuAttr，实际使用时请替换成真实参数
		},
	}

	UpdateOrderInfoRequestDataItemEXTMainOrderPaymentPreferentialInfoItemVal := openapi.UpdateOrderInfoRequestDataItemEXTMainOrderPaymentPreferentialInfoItem{
		Name:     "优惠券使用", // 文档中对应字段：Name，实际使用时请替换成真实参数
		Quantity: 1,       // 文档中对应字段：Quantity，实际使用时请替换成真实参数
		Value:    100,     // 文档中对应字段：Value，实际使用时请替换成真实参数
	}

	UpdateOrderInfoRequestDataItemEXTMainOrderPaymentPaymentInfoItemVal := openapi.UpdateOrderInfoRequestDataItemEXTMainOrderPaymentPaymentInfoItem{
		Name:     "运费", // 文档中对应字段：Name，实际使用时请替换成真实参数
		Quantity: 1,    // 文档中对应字段：Quantity，实际使用时请替换成真实参数
		Value:    100,  // 文档中对应字段：Value，实际使用时请替换成真实参数
	}

	UpdateOrderInfoRequestDataItemEXTMainOrderPaymentVal := openapi.UpdateOrderInfoRequestDataItemEXTMainOrderPayment{
		Amount:    2390,  // 文档中对应字段：Amount，实际使用时请替换成真实参数
		IsPayment: false, // 文档中对应字段：IsPayment，实际使用时请替换成真实参数
		Method:    1,     // 文档中对应字段：Method，实际使用时请替换成真实参数
		PaymentInfo: []openapi.UpdateOrderInfoRequestDataItemEXTMainOrderPaymentPaymentInfoItem{
			UpdateOrderInfoRequestDataItemEXTMainOrderPaymentPaymentInfoItemVal, // 文档中对应字段：PaymentInfo，实际使用时请替换成真实参数
		},
		PreferentialInfo: []openapi.UpdateOrderInfoRequestDataItemEXTMainOrderPaymentPreferentialInfoItem{
			UpdateOrderInfoRequestDataItemEXTMainOrderPaymentPreferentialInfoItemVal, // 文档中对应字段：PreferentialInfo，实际使用时请替换成真实参数
		},
		Time: 0, // 文档中对应字段：Time，实际使用时请替换成真实参数
	}

	UpdateOrderInfoRequestDataItemEXTMainOrderAppraiseVal := openapi.UpdateOrderInfoRequestDataItemEXTMainOrderAppraise{
		Status:     0,                                                                                                             // 文档中对应字段：Status，实际使用时请替换成真实参数
		SwanSchema: "baiduboxapp://swan/B8GF7AWvCSr18myIs81uqaoYz3pPCSY4/wjz/bdxd/order-detail/order-detail?orderId=430221766721", // 文档中对应字段：SwanSchema，实际使用时请替换成真实参数
	}

	UpdateOrderInfoRequestDataItemEXTMainOrderOrderDetailVal := openapi.UpdateOrderInfoRequestDataItemEXTMainOrderOrderDetail{
		Status:     2,                                                                                                             // 文档中对应字段：Status，实际使用时请替换成真实参数
		SwanSchema: "baiduboxapp://swan/B7GF2AWvCSr14myIs82uqaoYz2pPCSY6/wjz/bdxd/order-detail/order-detail?orderId=746647786177", // 文档中对应字段：SwanSchema，实际使用时请替换成真实参数
	}

	UpdateOrderInfoRequestDataItemEXTMainOrderVal := openapi.UpdateOrderInfoRequestDataItemEXTMainOrder{
		Appraise:    UpdateOrderInfoRequestDataItemEXTMainOrderAppraiseVal,    // 文档中对应字段：Appraise，实际使用时请替换成真实参数
		OrderDetail: UpdateOrderInfoRequestDataItemEXTMainOrderOrderDetailVal, // 文档中对应字段：OrderDetail，实际使用时请替换成真实参数
		Payment:     UpdateOrderInfoRequestDataItemEXTMainOrderPaymentVal,     // 文档中对应字段：Payment，实际使用时请替换成真实参数
		Products: []openapi.UpdateOrderInfoRequestDataItemEXTMainOrderProductsItem{
			UpdateOrderInfoRequestDataItemEXTMainOrderProductsItemVal, // 文档中对应字段：Products，实际使用时请替换成真实参数
		},
	}

	UpdateOrderInfoRequestDataItemEXTVal := openapi.UpdateOrderInfoRequestDataItemEXT{
		MainOrder: UpdateOrderInfoRequestDataItemEXTMainOrderVal, // 文档中对应字段：MainOrder，实际使用时请替换成真实参数
	}

	UpdateOrderInfoRequestDataItemVal := openapi.UpdateOrderInfoRequestDataItem{
		BizAPPID:   "WXF3pGOvo4TTGU4qCMMhEjvFBkF3bO4Z",   // 文档中对应字段：BizAPPID，实际使用时请替换成真实参数
		CateID:     1,                                    // 文档中对应字段：CateID，实际使用时请替换成真实参数
		EXT:        UpdateOrderInfoRequestDataItemEXTVal, // 文档中对应字段：EXT，实际使用时请替换成真实参数
		ResourceID: "1801025661622",                      // 文档中对应字段：ResourceID，实际使用时请替换成真实参数
		Status:     200,                                  // 文档中对应字段：Status，实际使用时请替换成真实参数
	}

	reqParams := &openapi.UpdateOrderInfoRequest{
		AccessToken: "22.306f00223a0740a7c8ad6c405666eb82.1143604.3177761583.547381-70532611", // 文档中对应字段：access_token，实际使用时请替换成真实参数
		OpenID:      "k82HEREQhWhWWB5WYqYT0ITUGX",                                             // 文档中对应字段：open_id，实际使用时请替换成真实参数
		SceneID:     "1183337360313",                                                          // 文档中对应字段：scene_id，实际使用时请替换成真实参数
		SceneType:   2,                                                                        // 文档中对应字段：scene_type，实际使用时请替换成真实参数
		PmAppKey:    "baiduboxapp",                                                            // 文档中对应字段：pm_app_key，实际使用时请替换成真实参数
		Data: []openapi.UpdateOrderInfoRequestDataItem{
			UpdateOrderInfoRequestDataItemVal,
		}, // 文档中对应字段：Data，实际使用时请替换成真实参数
	}

	resp, err := openapi.UpdateOrderInfo(reqParams)
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

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
	AddOrderInfoRequestDataItemEXTMainOrderProductsItemDetailPageVal := openapi.AddOrderInfoRequestDataItemEXTMainOrderProductsItemDetailPage{
		Status:     "2",                                                                                                           // 文档中对应字段：Status，实际使用时请替换成真实参数
		SwanSchema: "baiduboxapp://swan/B0GF5AWvCSr44myIs76uqaoYz2pPCSY5/wjz/bdxd/order-detail/order-detail?orderId=263208875771", // 文档中对应字段：SwanSchema，实际使用时请替换成真实参数
	}

	AddOrderInfoRequestDataItemEXTMainOrderProductsItemSkuAttrItemVal := openapi.AddOrderInfoRequestDataItemEXTMainOrderProductsItemSkuAttrItem{
		Name:  "四川大凉山丑苹果脆甜",  // 文档中对应字段：Name，实际使用时请替换成真实参数
		Value: "5斤小果25个左右偏小", // 文档中对应字段：Value，实际使用时请替换成真实参数
	}

	AddOrderInfoRequestDataItemEXTMainOrderProductsItemVal := openapi.AddOrderInfoRequestDataItemEXTMainOrderProductsItem{
		Desc:       "四川大凉山丑苹果脆甜:5斤小果25个左右偏小;",                                        // 文档中对应字段：Desc，实际使用时请替换成真实参数
		DetailPage: AddOrderInfoRequestDataItemEXTMainOrderProductsItemDetailPageVal, // 文档中对应字段：DetailPage，实际使用时请替换成真实参数
		ID:         "1336064258405",                                                  // 文档中对应字段：ID，实际使用时请替换成真实参数
		ImgList: []string{
			"<no value>", // 文档中对应字段：ImgList，实际使用时请替换成真实参数
		},
		Name:     "四川大凉山丑苹果脆甜红将军盐源丑苹果", // 文档中对应字段：Name，实际使用时请替换成真实参数
		PayPrice: 2390,                 // 文档中对应字段：PayPrice，实际使用时请替换成真实参数
		Price:    2390,                 // 文档中对应字段：Price，实际使用时请替换成真实参数
		Quantity: 1,                    // 文档中对应字段：Quantity，实际使用时请替换成真实参数
		SkuAttr: []openapi.AddOrderInfoRequestDataItemEXTMainOrderProductsItemSkuAttrItem{
			AddOrderInfoRequestDataItemEXTMainOrderProductsItemSkuAttrItemVal, // 文档中对应字段：SkuAttr，实际使用时请替换成真实参数
		},
	}

	AddOrderInfoRequestDataItemEXTMainOrderPaymentPreferentialInfoItemVal := openapi.AddOrderInfoRequestDataItemEXTMainOrderPaymentPreferentialInfoItem{
		Name:     "优惠券使用", // 文档中对应字段：Name，实际使用时请替换成真实参数
		Quantity: 1,       // 文档中对应字段：Quantity，实际使用时请替换成真实参数
		Value:    100,     // 文档中对应字段：Value，实际使用时请替换成真实参数
	}

	AddOrderInfoRequestDataItemEXTMainOrderPaymentPaymentInfoItemVal := openapi.AddOrderInfoRequestDataItemEXTMainOrderPaymentPaymentInfoItem{
		Name:     "运费", // 文档中对应字段：Name，实际使用时请替换成真实参数
		Quantity: 1,    // 文档中对应字段：Quantity，实际使用时请替换成真实参数
		Value:    100,  // 文档中对应字段：Value，实际使用时请替换成真实参数
	}

	AddOrderInfoRequestDataItemEXTMainOrderPaymentVal := openapi.AddOrderInfoRequestDataItemEXTMainOrderPayment{
		Amount:    2390,  // 文档中对应字段：Amount，实际使用时请替换成真实参数
		IsPayment: false, // 文档中对应字段：IsPayment，实际使用时请替换成真实参数
		Method:    1,     // 文档中对应字段：Method，实际使用时请替换成真实参数
		PaymentInfo: []openapi.AddOrderInfoRequestDataItemEXTMainOrderPaymentPaymentInfoItem{
			AddOrderInfoRequestDataItemEXTMainOrderPaymentPaymentInfoItemVal, // 文档中对应字段：PaymentInfo，实际使用时请替换成真实参数
		},
		PreferentialInfo: []openapi.AddOrderInfoRequestDataItemEXTMainOrderPaymentPreferentialInfoItem{
			AddOrderInfoRequestDataItemEXTMainOrderPaymentPreferentialInfoItemVal, // 文档中对应字段：PreferentialInfo，实际使用时请替换成真实参数
		},
		Time: 0, // 文档中对应字段：Time，实际使用时请替换成真实参数
	}

	AddOrderInfoRequestDataItemEXTMainOrderAppraiseVal := openapi.AddOrderInfoRequestDataItemEXTMainOrderAppraise{
		Status:     0,                                                                                                             // 文档中对应字段：Status，实际使用时请替换成真实参数
		SwanSchema: "baiduboxapp://swan/B1GF8AWvCSr30myIs72uqaoYz1pPCSY0/wjz/bdxd/order-detail/order-detail?orderId=867106411087", // 文档中对应字段：SwanSchema，实际使用时请替换成真实参数
	}

	AddOrderInfoRequestDataItemEXTMainOrderOrderDetailVal := openapi.AddOrderInfoRequestDataItemEXTMainOrderOrderDetail{
		Status:     0,                                                                                                             // 文档中对应字段：Status，实际使用时请替换成真实参数
		SwanSchema: "baiduboxapp://swan/B6GF7AWvCSr60myIs82uqaoYz5pPCSY6/wjz/bdxd/order-detail/order-detail?orderId=436884243764", // 文档中对应字段：SwanSchema，实际使用时请替换成真实参数
	}

	AddOrderInfoRequestDataItemEXTMainOrderVal := openapi.AddOrderInfoRequestDataItemEXTMainOrder{
		Appraise:    AddOrderInfoRequestDataItemEXTMainOrderAppraiseVal,    // 文档中对应字段：Appraise，实际使用时请替换成真实参数
		OrderDetail: AddOrderInfoRequestDataItemEXTMainOrderOrderDetailVal, // 文档中对应字段：OrderDetail，实际使用时请替换成真实参数
		Payment:     AddOrderInfoRequestDataItemEXTMainOrderPaymentVal,     // 文档中对应字段：Payment，实际使用时请替换成真实参数
		Products: []openapi.AddOrderInfoRequestDataItemEXTMainOrderProductsItem{
			AddOrderInfoRequestDataItemEXTMainOrderProductsItemVal, // 文档中对应字段：Products，实际使用时请替换成真实参数
		},
	}

	AddOrderInfoRequestDataItemEXTVal := openapi.AddOrderInfoRequestDataItemEXT{
		MainOrder: AddOrderInfoRequestDataItemEXTMainOrderVal, // 文档中对应字段：MainOrder，实际使用时请替换成真实参数
	}

	AddOrderInfoRequestDataItemVal := openapi.AddOrderInfoRequestDataItem{
		BizAPPID:   "WXF8pGOvo2TTGU3qCMMhEjvFBkF2bO8Z", // 文档中对应字段：BizAPPID，实际使用时请替换成真实参数
		CateID:     1,                                  // 文档中对应字段：CateID，实际使用时请替换成真实参数
		Ctime:      1233212343,                         // 文档中对应字段：Ctime，实际使用时请替换成真实参数
		EXT:        AddOrderInfoRequestDataItemEXTVal,  // 文档中对应字段：EXT，实际使用时请替换成真实参数
		Mtime:      1233212343,                         // 文档中对应字段：Mtime，实际使用时请替换成真实参数
		ResourceID: "1750884287254",                    // 文档中对应字段：ResourceID，实际使用时请替换成真实参数
		Status:     200,                                // 文档中对应字段：Status，实际使用时请替换成真实参数
		Title:      "test",                             // 文档中对应字段：Title，实际使用时请替换成真实参数
	}

	reqParams := &openapi.AddOrderInfoRequest{
		AccessToken: "24.808f42642a4855a6c6ad3c053451eb13.1521086.1735470025.856567-05072187", // 文档中对应字段：access_token，实际使用时请替换成真实参数
		OpenID:      "k33HEREQhWhWWB1WYqYT3ITUGX",                                             // 文档中对应字段：open_id，实际使用时请替换成真实参数
		SwanID:      "",                                                                       // 文档中对应字段：swan_id，实际使用时请替换成真实参数
		SceneID:     "1188213347002",                                                          // 文档中对应字段：scene_id，实际使用时请替换成真实参数
		SceneType:   2,                                                                        // 文档中对应字段：scene_type，实际使用时请替换成真实参数
		PmAppKey:    "baiduboxapp",                                                            // 文档中对应字段：pm_app_key，实际使用时请替换成真实参数
		Data: []openapi.AddOrderInfoRequestDataItem{
			AddOrderInfoRequestDataItemVal,
		}, // 文档中对应字段：Data，实际使用时请替换成真实参数
	}

	resp, err := openapi.AddOrderInfo(reqParams)
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

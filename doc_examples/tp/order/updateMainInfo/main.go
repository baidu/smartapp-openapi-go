// 本示例基于百度智能小程序服务端开发者 smartapp-openapi-go SDK
// 使用该示例需要首先下载该 SDK，如果是第一次使用，可使用以下命令拉取依赖
// go get github.com/baidu/smartapp-openapi-go
// 如使用过程中遇到问题，可以加入如流群：5702992，进行反馈咨询
package main

import (
	"fmt"

	"github.com/baidu/smartapp-openapi-go/openapi/tp/order"
	"github.com/baidu/smartapp-openapi-go/utils"
)

func main() {
	// 开发者在此设置请求参数，文档示例中的参数均为示例参数，实际参数请参考对应接口的文档上方的参数说明填写
	// 注意：代码示例中的参数字段基本是驼峰形式，而文档中的参数说明的参数字段基本是下划线形式
	// 如果开发者不想传非必需参数，可以将设置该参数的行注释
	UpdateMainInfoRequestDataItemEXTMainOrderProductsItemSkuAttrItemVal := order.UpdateMainInfoRequestDataItemEXTMainOrderProductsItemSkuAttrItem{
		Name:  "", // 文档中对应字段：Name，实际使用时请替换成真实参数
		Value: "", // 文档中对应字段：Value，实际使用时请替换成真实参数
	}

	UpdateMainInfoRequestDataItemEXTMainOrderProductsItemVal := order.UpdateMainInfoRequestDataItemEXTMainOrderProductsItem{
		Desc: "", // 文档中对应字段：Desc，实际使用时请替换成真实参数
		ID:   "", // 文档中对应字段：ID，实际使用时请替换成真实参数
		ImgList: []string{
			"<no value>", // 文档中对应字段：ImgList，实际使用时请替换成真实参数
		},
		Name:     "", // 文档中对应字段：Name，实际使用时请替换成真实参数
		PayPrice: 0,  // 文档中对应字段：PayPrice，实际使用时请替换成真实参数
		Price:    0,  // 文档中对应字段：Price，实际使用时请替换成真实参数
		Quantity: 0,  // 文档中对应字段：Quantity，实际使用时请替换成真实参数
		SkuAttr: []order.UpdateMainInfoRequestDataItemEXTMainOrderProductsItemSkuAttrItem{
			UpdateMainInfoRequestDataItemEXTMainOrderProductsItemSkuAttrItemVal, // 文档中对应字段：SkuAttr，实际使用时请替换成真实参数
		},
	}

	UpdateMainInfoRequestDataItemEXTMainOrderPaymentPreferentialInfoItemVal := order.UpdateMainInfoRequestDataItemEXTMainOrderPaymentPreferentialInfoItem{
		Name:     "", // 文档中对应字段：Name，实际使用时请替换成真实参数
		Quantity: 0,  // 文档中对应字段：Quantity，实际使用时请替换成真实参数
		Value:    0,  // 文档中对应字段：Value，实际使用时请替换成真实参数
	}

	UpdateMainInfoRequestDataItemEXTMainOrderPaymentPaymentInfoItemVal := order.UpdateMainInfoRequestDataItemEXTMainOrderPaymentPaymentInfoItem{
		Name:     "", // 文档中对应字段：Name，实际使用时请替换成真实参数
		Quantity: 0,  // 文档中对应字段：Quantity，实际使用时请替换成真实参数
		Value:    0,  // 文档中对应字段：Value，实际使用时请替换成真实参数
	}

	UpdateMainInfoRequestDataItemEXTMainOrderPaymentVal := order.UpdateMainInfoRequestDataItemEXTMainOrderPayment{
		Amount:    0,     // 文档中对应字段：Amount，实际使用时请替换成真实参数
		IsPayment: false, // 文档中对应字段：IsPayment，实际使用时请替换成真实参数
		Method:    0,     // 文档中对应字段：Method，实际使用时请替换成真实参数
		PaymentInfo: []order.UpdateMainInfoRequestDataItemEXTMainOrderPaymentPaymentInfoItem{
			UpdateMainInfoRequestDataItemEXTMainOrderPaymentPaymentInfoItemVal, // 文档中对应字段：PaymentInfo，实际使用时请替换成真实参数
		},
		PreferentialInfo: []order.UpdateMainInfoRequestDataItemEXTMainOrderPaymentPreferentialInfoItem{
			UpdateMainInfoRequestDataItemEXTMainOrderPaymentPreferentialInfoItemVal, // 文档中对应字段：PreferentialInfo，实际使用时请替换成真实参数
		},
		Time: 0, // 文档中对应字段：Time，实际使用时请替换成真实参数
	}

	UpdateMainInfoRequestDataItemEXTMainOrderAppraiseVal := order.UpdateMainInfoRequestDataItemEXTMainOrderAppraise{
		Name:       "", // 文档中对应字段：Name，实际使用时请替换成真实参数
		Status:     0,  // 文档中对应字段：Status，实际使用时请替换成真实参数
		SwanSchema: "", // 文档中对应字段：SwanSchema，实际使用时请替换成真实参数
	}

	UpdateMainInfoRequestDataItemEXTMainOrderOrderDetailVal := order.UpdateMainInfoRequestDataItemEXTMainOrderOrderDetail{
		Name:       "", // 文档中对应字段：Name，实际使用时请替换成真实参数
		Status:     0,  // 文档中对应字段：Status，实际使用时请替换成真实参数
		SwanSchema: "", // 文档中对应字段：SwanSchema，实际使用时请替换成真实参数
	}

	UpdateMainInfoRequestDataItemEXTMainOrderVal := order.UpdateMainInfoRequestDataItemEXTMainOrder{
		Appraise:    UpdateMainInfoRequestDataItemEXTMainOrderAppraiseVal,    // 文档中对应字段：Appraise，实际使用时请替换成真实参数
		OrderDetail: UpdateMainInfoRequestDataItemEXTMainOrderOrderDetailVal, // 文档中对应字段：OrderDetail，实际使用时请替换成真实参数
		Payment:     UpdateMainInfoRequestDataItemEXTMainOrderPaymentVal,     // 文档中对应字段：Payment，实际使用时请替换成真实参数
		Products: []order.UpdateMainInfoRequestDataItemEXTMainOrderProductsItem{
			UpdateMainInfoRequestDataItemEXTMainOrderProductsItemVal, // 文档中对应字段：Products，实际使用时请替换成真实参数
		},
	}

	UpdateMainInfoRequestDataItemEXTVal := order.UpdateMainInfoRequestDataItemEXT{
		MainOrder: UpdateMainInfoRequestDataItemEXTMainOrderVal, // 文档中对应字段：MainOrder，实际使用时请替换成真实参数
	}

	UpdateMainInfoRequestDataItemVal := order.UpdateMainInfoRequestDataItem{
		BizAPPID:   "",                                  // 文档中对应字段：BizAPPID，实际使用时请替换成真实参数
		CateID:     0,                                   // 文档中对应字段：CateID，实际使用时请替换成真实参数
		EXT:        UpdateMainInfoRequestDataItemEXTVal, // 文档中对应字段：EXT，实际使用时请替换成真实参数
		ResourceID: "",                                  // 文档中对应字段：ResourceID，实际使用时请替换成真实参数
		Status:     0,                                   // 文档中对应字段：Status，实际使用时请替换成真实参数
	}

	reqParams := &order.UpdateMainInfoRequest{
		AccessToken: "#token", // 文档中对应字段：access_token，实际使用时请替换成真实参数
		OpenID:      "128",    // 文档中对应字段：open_id，实际使用时请替换成真实参数
		SceneID:     "100",    // 文档中对应字段：scene_id，实际使用时请替换成真实参数
		SceneType:   2,        // 文档中对应字段：scene_type，实际使用时请替换成真实参数
		Data: []order.UpdateMainInfoRequestDataItem{
			UpdateMainInfoRequestDataItemVal,
		}, // 文档中对应字段：Data，实际使用时请替换成真实参数
	}

	resp, err := order.UpdateMainInfo(reqParams)
	if err != nil {
		if _, ok := err.(*utils.OpenAPIError); ok {
			// openapi error
			// 可能是 access_token 无效，可以尝试重新生成 access_token
			fmt.Println("openapi error, ", err)
		} else if _, ok := err.(*utils.APIError); ok {
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

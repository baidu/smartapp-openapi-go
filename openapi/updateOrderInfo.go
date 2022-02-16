package openapi

import (
	"encoding/json"
	"fmt"
)

// UpdateOrderInfoRequest 请求结构体
type UpdateOrderInfoRequest struct {
	AccessToken string                           // 接口调用凭证
	OpenID      string                           // 用户 openId
	SceneID     string                           // 百度收银台分配的平台订单 ID，通知支付状态接口返回的 orderId
	SceneType   int64                            // 支付场景类型，开发者请默认传 2
	PmAppKey    string                           // 调起百度收银台的支付服务 appKey
	Data        []UpdateOrderInfoRequestDataItem // 请求数据
}

type UpdateOrderInfoRequestDataItemEXTMainOrderOrderDetail struct {
	Status     int64  `json:"Status"`     // 默认传 2
	SwanSchema string `json:"SwanSchema"` // 订单详情页的跳转地址，用以小程序跳转 Scheme
}
type UpdateOrderInfoRequestDataItemEXTMainOrderProductsItemDetailPage struct {
	Status     int64  `json:"Status"`     // 默认传 2
	SwanSchema string `json:"SwanSchema"` // 商品详情页的跳转地址，用以小程序跳转 Scheme
}
type UpdateOrderInfoRequestDataItemEXTMainOrderProductsItemSkuAttrItem struct {
	Name  string `json:"Name"`  // 规格名称，例如“颜色”或“尺寸”
	Value string `json:"Value"` // 规格值
}
type UpdateOrderInfoRequestDataItemEXTMainOrderProductsItem struct {
	Desc       string                                                              `json:"Desc"`       // 商品简述
	DetailPage UpdateOrderInfoRequestDataItemEXTMainOrderProductsItemDetailPage    `json:"DetailPage"` // 商品详情的跳转的跳转结构
	ID         string                                                              `json:"ID"`         // 商品 ID ，开发者的唯一商品 ID
	ImgList    []string                                                            `json:"ImgList"`    // 商品预览，值为预览图 URL 地址，最多 5 张
	Name       string                                                              `json:"Name"`       // 商品名字
	PayPrice   int64                                                               `json:"PayPrice"`   // 实付价（单位：分），即100代表1元
	Price      int64                                                               `json:"Price"`      // 本商品原价（单位：分），即100代表1元
	Quantity   int64                                                               `json:"Quantity"`   // 本商品的交易数量
	SkuAttr    []UpdateOrderInfoRequestDataItemEXTMainOrderProductsItemSkuAttrItem `json:"SkuAttr"`    // 商品规格，最多 400 个
}
type UpdateOrderInfoRequestDataItemEXTMainOrderPaymentPreferentialInfoItem struct {
	Name     string `json:"Name"`     // 展示名称
	Quantity int64  `json:"Quantity"` // 数量
	Value    int64  `json:"Value"`    // 合计金额（单位：分），即100为1元
}
type UpdateOrderInfoRequestDataItemEXTMainOrderPaymentPaymentInfoItem struct {
	Name     string `json:"Name"`     // 展示名称
	Quantity int64  `json:"Quantity"` // 数量
	Value    int64  `json:"Value"`    // 合计金额（单位：分），即100为1元
}
type UpdateOrderInfoRequestDataItemEXTMainOrderPayment struct {
	Amount           int64                                                                   `json:"Amount"`           // 实付金额（单位：分），即100为1元
	IsPayment        bool                                                                    `json:"IsPayment"`        // 是否已付款
	Method           int64                                                                   `json:"Method"`           // 付款方式，1（在线付），2（货到付款）
	PaymentInfo      []UpdateOrderInfoRequestDataItemEXTMainOrderPaymentPaymentInfoItem      `json:"PaymentInfo"`      // 其他付款信息，如运费、保险等
	PreferentialInfo []UpdateOrderInfoRequestDataItemEXTMainOrderPaymentPreferentialInfoItem `json:"PreferentialInfo"` // 优惠券信息
	Time             int64                                                                   `json:"Time"`             // 付款时间（单位：秒）
}
type UpdateOrderInfoRequestDataItemEXTMainOrderAppraise struct {
	Status     int64  `json:"Status"`     // 0（不可评价状态或已评价状态）、2（待评价状态，允许跳转)
	SwanSchema string `json:"SwanSchema"` // 评价页的跳转地址，用以小程序跳转 Scheme
}
type UpdateOrderInfoRequestDataItemEXTMainOrder struct {
	Appraise    UpdateOrderInfoRequestDataItemEXTMainOrderAppraise       `json:"Appraise"`    // 待评价状态订单的评价页结构，仅订单为可评价状态，且还未进行评价时提供该信息
	OrderDetail UpdateOrderInfoRequestDataItemEXTMainOrderOrderDetail    `json:"OrderDetail"` // 订单详情页的信息
	Payment     UpdateOrderInfoRequestDataItemEXTMainOrderPayment        `json:"Payment"`     // 支付信息
	Products    []UpdateOrderInfoRequestDataItemEXTMainOrderProductsItem `json:"Products"`    // 数组，商品信息列表，若商品只有 1 个则数组长度为 1
}
type UpdateOrderInfoRequestDataItemEXT struct {
	MainOrder UpdateOrderInfoRequestDataItemEXTMainOrder `json:"MainOrder"` // 主订单信息（购买商品订单）
}
type UpdateOrderInfoRequestDataItem struct {
	BizAPPID   string                            `json:"BizAPPID"`   // 小程序 AppKey
	CateID     int64                             `json:"CateID"`     // 订单种类：1（实物）、2（虚拟物品）、5（快递服务类）、6（快递服务类无金额订单）、10（上门服务类）、11（上门服务类无金额订单）、15（酒店类）、20（票务类）、25（打车类）、26（打车类无金额订单）
	EXT        UpdateOrderInfoRequestDataItemEXT `json:"EXT"`        // 扩展信息
	ResourceID string                            `json:"ResourceID"` // 开发者接入的唯一订单 ID
	Status     int64                             `json:"Status"`     // 订单状态，其值根据CateID不同有不同的定义。CateID = 1 实物订单、CateID = 2 虚拟物品订单、CateID = 5 快递服务类订单、CateID = 6 快递服务类无金额订单、CateID = 10 上门服务类订单、CateID = 11 上门服务类无金额订单、CateID = 15 酒店类订单、CateID = 20 出行票务类订单、CateID = 25 打车类订单、CateID = 26 打车类无金额订单
}

// 响应结构体

type UpdateOrderInfoResponsedataItem struct {
	BizAppID     string `json:"biz_app_id"`    // POST 请求参数中 BizAPPID
	CateID       string `json:"cate_id"`       // POST 请求参数中 CateID
	ResourceID   string `json:"resource_id"`   // POST 请求参数中 ResourceID
	RowsAffected string `json:"rows_affected"` // 请求受影响行数（即请求是否成功， 0 为失败，非 0 为成功）
}

type UpdateOrderInfoResponse struct {
	Data      []UpdateOrderInfoResponsedataItem `json:"data"`       // 响应对象
	Errno     int64                             `json:"errno"`      // 错误码
	ErrMsg    string                            `json:"msg"`        // 错误信息
	ErrorCode int64                             `json:"error_code"` // openapi 错误码
	ErrorMsg  string                            `json:"error_msg"`  // openapi 错误信息
}

// UpdateOrderInfo
func UpdateOrderInfo(params *UpdateOrderInfoRequest) ([]UpdateOrderInfoResponsedataItem, error) {
	var (
		err        error
		defaultRet []UpdateOrderInfoResponsedataItem
	)
	respData := &UpdateOrderInfoResponse{}

	client := NewHTTPClient().
		SetContentType(ContentTypeJSON).
		SetConverterType(ConverterTypeJSON).
		SetMethod("POST").
		SetScheme(SCHEME).
		SetHost(OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/ordercenter/app/update/main/info")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("open_id", params.OpenID)
	client.AddGetParam("scene_id", params.SceneID)
	client.AddGetParam("scene_type", fmt.Sprintf("%v", params.SceneType))
	client.AddGetParam("pm_app_key", params.PmAppKey)
	client.AddGetParam("sp_sdk_ver", SDKVERSION)
	client.AddGetParam("sp_sdk_lang", SDKLANG)
	postData := map[string]interface{}{
		"Data": params.Data,
	}
	bts, err := json.Marshal(postData)
	if err != nil {
		return defaultRet, err
	}
	client.SetBody(bts)

	err = client.Do()
	if err != nil {
		return defaultRet, err
	}
	err = client.Convert(respData)
	if err != nil {
		return defaultRet, err
	}
	if respData.ErrorCode != 0 {
		return defaultRet, &OpenAPIError{respData.ErrorCode, respData.ErrorMsg, respData}
	}
	if respData.Errno != 0 {
		return defaultRet, &APIError{respData.Errno, respData.ErrMsg, respData}
	}

	return respData.Data, nil
}

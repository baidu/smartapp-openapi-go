package order

import (
	"encoding/json"
	"fmt"

	"github.com/baidu/smartapp-openapi-go/utils"
)

// UpdateMainInfoRequest 请求结构体
type UpdateMainInfoRequest struct {
	AccessToken string                          // 授权小程序的接口调用凭据
	OpenID      string                          // 用户 openId
	SceneID     string                          // 百度收银台分配的平台订单 ID ，通知支付状态接口返回的 orderId
	SceneType   int64                           // 支付场景类型，开发者请默认传 2
	Data        []UpdateMainInfoRequestDataItem //
}

type UpdateMainInfoRequestDataItemEXTMainOrderOrderDetail struct {
	Name       string `json:"Name"`       //
	Status     int64  `json:"Status"`     //
	SwanSchema string `json:"SwanSchema"` //
}
type UpdateMainInfoRequestDataItemEXTMainOrderProductsItemSkuAttrItem struct {
	Name  string `json:"Name"`  //
	Value string `json:"Value"` //
}
type UpdateMainInfoRequestDataItemEXTMainOrderProductsItem struct {
	Desc     string                                                             `json:"Desc"`     // 商品详情
	ID       string                                                             `json:"ID"`       // 商品ID
	ImgList  []string                                                           `json:"ImgList"`  // 商品图片地址
	Name     string                                                             `json:"Name"`     // 商品名称
	PayPrice int64                                                              `json:"PayPrice"` // 实付价格,单位分。
	Price    int64                                                              `json:"Price"`    // 商品原价,单位分。
	Quantity int64                                                              `json:"Quantity"` // 商品数量
	SkuAttr  []UpdateMainInfoRequestDataItemEXTMainOrderProductsItemSkuAttrItem `json:"SkuAttr"`  // 商品SKU属性
}
type UpdateMainInfoRequestDataItemEXTMainOrderPaymentPreferentialInfoItem struct {
	Name     string `json:"Name"`     // 名称
	Quantity int64  `json:"Quantity"` // 数量
	Value    int64  `json:"Value"`    // 优惠金额，单位分
}
type UpdateMainInfoRequestDataItemEXTMainOrderPaymentPaymentInfoItem struct {
	Name     string `json:"Name"`     // 展示名称
	Quantity int64  `json:"Quantity"` // 数量
	Value    int64  `json:"Value"`    // 付款金额，单位分
}
type UpdateMainInfoRequestDataItemEXTMainOrderPayment struct {
	Amount           int64                                                                  `json:"Amount"`           // 合计金额，单位分
	IsPayment        bool                                                                   `json:"IsPayment"`        // 是否支付
	Method           int64                                                                  `json:"Method"`           // 支付方式
	PaymentInfo      []UpdateMainInfoRequestDataItemEXTMainOrderPaymentPaymentInfoItem      `json:"PaymentInfo"`      // 付款信息
	PreferentialInfo []UpdateMainInfoRequestDataItemEXTMainOrderPaymentPreferentialInfoItem `json:"PreferentialInfo"` // 优惠信息
	Time             int64                                                                  `json:"Time"`             // 付款时间，时间戳
}
type UpdateMainInfoRequestDataItemEXTMainOrderAppraise struct {
	Name       string `json:"Name"`       //
	Status     int64  `json:"Status"`     //
	SwanSchema string `json:"SwanSchema"` //
}
type UpdateMainInfoRequestDataItemEXTMainOrder struct {
	Appraise    UpdateMainInfoRequestDataItemEXTMainOrderAppraise       `json:"Appraise"`    // 订单评价跳转
	OrderDetail UpdateMainInfoRequestDataItemEXTMainOrderOrderDetail    `json:"OrderDetail"` // 订单详情跳转
	Payment     UpdateMainInfoRequestDataItemEXTMainOrderPayment        `json:"Payment"`     // 支付信息
	Products    []UpdateMainInfoRequestDataItemEXTMainOrderProductsItem `json:"Products"`    // 商品信息
}
type UpdateMainInfoRequestDataItemEXT struct {
	MainOrder UpdateMainInfoRequestDataItemEXTMainOrder `json:"MainOrder"` // 订单信息
}
type UpdateMainInfoRequestDataItem struct {
	BizAPPID   string                           `json:"BizAPPID"`   // 小程序的key
	CateID     int64                            `json:"CateID"`     // 2:订单种类-虚拟物品
	EXT        UpdateMainInfoRequestDataItemEXT `json:"EXT"`        // 拓展字段 此处以订单为例
	ResourceID string                           `json:"ResourceID"` // 开发者接入的唯一订单ID
	Status     int64                            `json:"Status"`     // 200:订单状态-已完成交易
}

// 响应结构体

type UpdateMainInfoResponsedataItem struct {
	BizAppID     string `json:"biz_app_id"`    // POST请求参数中BizAPPID
	CateID       string `json:"cate_id"`       // POST请求参数中CateID
	ResourceID   string `json:"resource_id"`   // POST请求参数中ResourceID
	RowsAffected string `json:"rows_affected"` // 请求受影响行数(即请求是否成功 0为失败 非0为成功)
}

type UpdateMainInfoResponse struct {
	Data      []UpdateMainInfoResponsedataItem `json:"data"`       // 响应参数
	Errno     int64                            `json:"errno"`      // 状态码
	ErrMsg    string                           `json:"msg"`        // 错误信息
	ErrorCode int64                            `json:"error_code"` // openapi 错误码
	ErrorMsg  string                           `json:"error_msg"`  // openapi 错误信息
}

// UpdateMainInfo
func UpdateMainInfo(params *UpdateMainInfoRequest) ([]UpdateMainInfoResponsedataItem, error) {
	var (
		err        error
		defaultRet []UpdateMainInfoResponsedataItem
	)
	respData := &UpdateMainInfoResponse{}

	client := utils.NewHTTPClient().
		SetContentType(utils.ContentTypeJSON).
		SetConverterType(utils.ConverterTypeJSON).
		SetMethod("POST").
		SetScheme(utils.SCHEME).
		SetHost(utils.OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/ordercenter/update/main/info")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("open_id", params.OpenID)
	client.AddGetParam("scene_id", params.SceneID)
	client.AddGetParam("scene_type", fmt.Sprintf("%v", params.SceneType))
	client.AddGetParam("sp_sdk_ver", utils.SDKVERSION)
	client.AddGetParam("sp_sdk_lang", utils.SDKLANG)
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
		return defaultRet, &utils.OpenAPIError{respData.ErrorCode, respData.ErrorMsg, respData}
	}

	if respData.Errno != 0 {
		return defaultRet, &utils.APIError{respData.Errno, respData.ErrMsg, respData}
	}
	return respData.Data, nil
}

package order

import (
	"encoding/json"
	"fmt"

	"github.com/baidu/smartapp-openapi-go/utils"
)

// AddMainInfoRequest 请求结构体
type AddMainInfoRequest struct {
	AccessToken string                       // 授权小程序的接口调用凭据
	OpenID      string                       // 用户 openId
	SceneID     string                       // 百度收银台分配的平台订单 ID ，通知支付状态接口返回的 orderId
	SceneType   int64                        // 支付场景类型，开发者请默认传 2
	Data        []AddMainInfoRequestDataItem //
}

type AddMainInfoRequestDataItemEXTMainOrderProductsItemSkuAttrItem struct {
	Name  string `json:"Name"`  //
	Value string `json:"Value"` //
}
type AddMainInfoRequestDataItemEXTMainOrderProductsItem struct {
	Desc     string                                                          `json:"Desc"`     // 商品详情
	ID       string                                                          `json:"ID"`       // 商品ID
	ImgList  []string                                                        `json:"ImgList"`  // 商品图片地址
	Name     string                                                          `json:"Name"`     // 商品名称
	PayPrice int64                                                           `json:"PayPrice"` // 实付价格,单位分。
	Price    int64                                                           `json:"Price"`    // 商品原价,单位分。
	Quantity int64                                                           `json:"Quantity"` // 商品数量
	SkuAttr  []AddMainInfoRequestDataItemEXTMainOrderProductsItemSkuAttrItem `json:"SkuAttr"`  // 商品SKU属性
}
type AddMainInfoRequestDataItemEXTMainOrderPaymentPreferentialInfoItem struct {
	Name     string `json:"Name"`     // 名称
	Quantity int64  `json:"Quantity"` // 数量
	Value    int64  `json:"Value"`    // 优惠金额，单位分
}
type AddMainInfoRequestDataItemEXTMainOrderPaymentPaymentInfoItem struct {
	Name     string `json:"Name"`     // 展示名称
	Quantity int64  `json:"Quantity"` // 数量
	Value    int64  `json:"Value"`    // 付款金额，单位分
}
type AddMainInfoRequestDataItemEXTMainOrderPayment struct {
	Amount           int64                                                               `json:"Amount"`           // 合计金额，单位分
	IsPayment        bool                                                                `json:"IsPayment"`        // 是否支付
	Method           int64                                                               `json:"Method"`           // 支付方式
	PaymentInfo      []AddMainInfoRequestDataItemEXTMainOrderPaymentPaymentInfoItem      `json:"PaymentInfo"`      // 付款信息
	PreferentialInfo []AddMainInfoRequestDataItemEXTMainOrderPaymentPreferentialInfoItem `json:"PreferentialInfo"` // 优惠信息
	Time             int64                                                               `json:"Time"`             // 付款时间，时间戳
}
type AddMainInfoRequestDataItemEXTMainOrderAppraise struct {
	Name       string `json:"Name"`       //
	Status     int64  `json:"Status"`     // 订单评价跳转
	SwanSchema string `json:"SwanSchema"` //
}
type AddMainInfoRequestDataItemEXTMainOrderOrderDetail struct {
	Name       string `json:"Name"`       //
	Status     int64  `json:"Status"`     //
	SwanSchema string `json:"SwanSchema"` //
}
type AddMainInfoRequestDataItemEXTMainOrder struct {
	Appraise    AddMainInfoRequestDataItemEXTMainOrderAppraise       `json:"Appraise"`    //
	OrderDetail AddMainInfoRequestDataItemEXTMainOrderOrderDetail    `json:"OrderDetail"` // 订单详情跳转
	Payment     AddMainInfoRequestDataItemEXTMainOrderPayment        `json:"Payment"`     // 支付信息
	Products    []AddMainInfoRequestDataItemEXTMainOrderProductsItem `json:"Products"`    // 商品信息
}
type AddMainInfoRequestDataItemEXTSubsOrderItemsItemRefundProductItem struct {
	Amount   int64  `json:"Amount"`   // 应退金额,单位分
	ID       string `json:"ID"`       // 商品ID
	Quantity int64  `json:"Quantity"` // 商品退款/商品退货 数量
}
type AddMainInfoRequestDataItemEXTSubsOrderItemsItemRefund struct {
	Amount  int64                                                              `json:"Amount"`  // 退款总金额
	Product []AddMainInfoRequestDataItemEXTSubsOrderItemsItemRefundProductItem `json:"Product"` // 退款/退货商品
}
type AddMainInfoRequestDataItemEXTSubsOrderItemsItemOrderDetail struct {
	AndroidSchema string `json:"AndroidSchema"` //
	IPhoneSchema  string `json:"IPhoneSchema"`  //
	Name          string `json:"Name"`          //
	Status        int64  `json:"Status"`        //
	SwanSchema    string `json:"SwanSchema"`    //
}
type AddMainInfoRequestDataItemEXTSubsOrderItemsItem struct {
	CTime       int64                                                      `json:"CTime"`       // 售后订单创建时间,时间戳
	MTime       int64                                                      `json:"MTime"`       // 售后订单修改时间,时间戳
	OrderDetail AddMainInfoRequestDataItemEXTSubsOrderItemsItemOrderDetail `json:"OrderDetail"` // 退款退货订单详情跳转
	OrderType   int64                                                      `json:"OrderType"`   // 退款订单类型
	Refund      AddMainInfoRequestDataItemEXTSubsOrderItemsItemRefund      `json:"Refund"`      // 商品 退款／退货 信息
	SubOrderID  string                                                     `json:"SubOrderID"`  // 售后订单ID
	SubStatus   int64                                                      `json:"SubStatus"`   // 自订单状态,枚举参照 【退换货枚举值】
}
type AddMainInfoRequestDataItemEXTSubsOrder struct {
	Items  []AddMainInfoRequestDataItemEXTSubsOrderItemsItem `json:"Items"`  //
	Status int64                                             `json:"Status"` //
}
type AddMainInfoRequestDataItemEXT struct {
	MainOrder AddMainInfoRequestDataItemEXTMainOrder `json:"MainOrder"` // 订单信息
	SubsOrder AddMainInfoRequestDataItemEXTSubsOrder `json:"SubsOrder"` // 售后订单信息，若该订单发生退款/售后，需新增同步其售后订单的售后信息状态
}
type AddMainInfoRequestDataItem struct {
	BizAPPID   string                        `json:"BizAPPID"`   // 小程序AppKey
	CateID     int64                         `json:"CateID"`     // 1:订单种类-实物商品
	Ctime      int64                         `json:"Ctime"`      // 订单创建时间
	EXT        AddMainInfoRequestDataItemEXT `json:"EXT"`        // 拓展字段
	Mtime      int64                         `json:"Mtime"`      // 订单最后修改时间
	ResourceID string                        `json:"ResourceID"` // 开发者接入的唯一订单ID
	Status     int64                         `json:"Status"`     // 200:订单状态-已完成交易
	Title      string                        `json:"Title"`      // 订单名称
}

// 响应结构体

type AddMainInfoResponsedataItem struct {
	BizAppID     string `json:"biz_app_id"`    // POST请求参数中BizAPPID
	CateID       string `json:"cate_id"`       // POST请求参数中CateID
	ResourceID   string `json:"resource_id"`   // POST请求参数中ResourceID
	RowsAffected string `json:"rows_affected"` // 请求受影响行数(即请求是否成功 0为失败 非0为成功)
}

type AddMainInfoResponse struct {
	Data      []AddMainInfoResponsedataItem `json:"data"`       // 响应参数
	Errno     int64                         `json:"errno"`      // 状态码
	ErrMsg    string                        `json:"msg"`        // 错误信息
	ErrorCode int64                         `json:"error_code"` // openapi 错误码
	ErrorMsg  string                        `json:"error_msg"`  // openapi 错误信息
}

// AddMainInfo
func AddMainInfo(params *AddMainInfoRequest) ([]AddMainInfoResponsedataItem, error) {
	var (
		err        error
		defaultRet []AddMainInfoResponsedataItem
	)
	respData := &AddMainInfoResponse{}

	client := utils.NewHTTPClient().
		SetContentType(utils.ContentTypeJSON).
		SetConverterType(utils.ConverterTypeJSON).
		SetMethod("POST").
		SetScheme(utils.SCHEME).
		SetHost(utils.OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/ordercenter/add/main/info")
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

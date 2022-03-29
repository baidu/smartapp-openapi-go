package order

import (
	"encoding/json"
	"fmt"

	"github.com/baidu/smartapp-openapi-go/utils"
)

// AppendSubInfoRequest 请求结构体
type AppendSubInfoRequest struct {
	AccessToken string                         // 授权小程序的接口调用凭据
	OpenID      string                         // 用户openId
	SceneID     string                         // 百度收银台分配的平台订单ID，通知支付状态接口返回的orderId
	SceneType   int64                          // 支付场景类型，开发者请默认传2
	Data        []AppendSubInfoRequestDataItem //
}

type AppendSubInfoRequestDataItemEXTSubsOrderItemsItemOrderDetail struct {
	Name       string `json:"Name"`       //
	Status     int64  `json:"Status"`     //
	SwanSchema string `json:"SwanSchema"` //
}
type AppendSubInfoRequestDataItemEXTSubsOrderItemsItemRefundProductItem struct {
	Amount   int64  `json:"Amount"`   // 应退金额,单位分
	ID       string `json:"ID"`       // 商品ID
	Quantity int64  `json:"Quantity"` // 商品退款/商品退货 数量
}
type AppendSubInfoRequestDataItemEXTSubsOrderItemsItemRefund struct {
	Amount  int64                                                                `json:"Amount"`  // 退款总金额
	Product []AppendSubInfoRequestDataItemEXTSubsOrderItemsItemRefundProductItem `json:"Product"` // 退款/退货商品
}
type AppendSubInfoRequestDataItemEXTSubsOrderItemsItem struct {
	CTime       int64                                                        `json:"CTime"`       // 售后订单创建时间,时间戳
	MTime       int64                                                        `json:"MTime"`       // 售后订单修改时间,时间戳
	OrderDetail AppendSubInfoRequestDataItemEXTSubsOrderItemsItemOrderDetail `json:"OrderDetail"` // 退款退货订单详情跳转
	OrderType   int64                                                        `json:"OrderType"`   // 退款订单类型
	Refund      AppendSubInfoRequestDataItemEXTSubsOrderItemsItemRefund      `json:"Refund"`      // 商品 退款／退货 信息
	SubOrderID  string                                                       `json:"SubOrderID"`  // 售后订单ID
	SubStatus   int64                                                        `json:"SubStatus"`   // 自订单状态,枚举参照 【退换货枚举值】
}
type AppendSubInfoRequestDataItemEXTSubsOrder struct {
	Items []AppendSubInfoRequestDataItemEXTSubsOrderItemsItem `json:"Items"` //
}
type AppendSubInfoRequestDataItemEXT struct {
	SubsOrder AppendSubInfoRequestDataItemEXTSubsOrder `json:"SubsOrder"` //  售后订单信息
}
type AppendSubInfoRequestDataItem struct {
	BizAPPID   string                          `json:"BizAPPID"`   // 小程序的appKey
	CateID     int64                           `json:"CateID"`     // 2:订单种类-虚拟物品
	EXT        AppendSubInfoRequestDataItemEXT `json:"EXT"`        // 拓展字段 根据资产的不同其结构也不固定 此处以订单为例
	ResourceID string                          `json:"ResourceID"` // 开发者接入的唯一订单ID
}

// 响应结构体

type AppendSubInfoResponsedataItem struct {
	BizAppID     string `json:"biz_app_id"`    // POST请求参数中BizAPPID
	CateID       string `json:"cate_id"`       // POST请求参数中CateID
	ResourceID   string `json:"resource_id"`   // POST请求参数中ResourceID
	RowsAffected string `json:"rows_affected"` // 请求受影响行数(即请求是否成功 0为失败 非0为成功)
}

type AppendSubInfoResponse struct {
	Data      []AppendSubInfoResponsedataItem `json:"data"`       // 响应参数
	Errno     int64                           `json:"errno"`      // 状态码
	ErrMsg    string                          `json:"msg"`        // 错误信息
	ErrorCode int64                           `json:"error_code"` // openapi 错误码
	ErrorMsg  string                          `json:"error_msg"`  // openapi 错误信息
}

// AppendSubInfo
func AppendSubInfo(params *AppendSubInfoRequest) ([]AppendSubInfoResponsedataItem, error) {
	var (
		err        error
		defaultRet []AppendSubInfoResponsedataItem
	)
	respData := &AppendSubInfoResponse{}

	client := utils.NewHTTPClient().
		SetContentType(utils.ContentTypeJSON).
		SetConverterType(utils.ConverterTypeJSON).
		SetMethod("POST").
		SetScheme(utils.SCHEME).
		SetHost(utils.OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/ordercenter/append/sub/info")
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

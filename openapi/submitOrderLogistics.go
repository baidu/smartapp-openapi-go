package openapi

import (
	"encoding/json"
	"fmt"
)

// SubmitOrderLogisticsRequest 请求结构体
type SubmitOrderLogisticsRequest struct {
	AccessToken string                                // 接口调用凭证
	OpenID      string                                // 用户 openId
	SceneID     string                                // 百度收银台分配的平台订单 ID，通知支付状态接口返回的 orderId
	SceneType   int64                                 // 支付场景类型，开发者请默认传 2
	PmAppKey    string                                // 调起百度收银台的支付服务 appKey
	Data        []SubmitOrderLogisticsRequestDataItem // 请求数据
}

type SubmitOrderLogisticsRequestDataItemEXTMainOrderExpress struct {
	Code   string `json:"Code"`   // 快递公司对应的编号，详情请参考快递公司信息码表
	ID     string `json:"ID"`     // 快递单号
	Name   string `json:"Name"`   // 快递公司名称
	Status int64  `json:"Status"` // 开发者默认传 0
	Type   int64  `json:"Type"`   // 快递类型，1：商家给用户发货，2：用户给商家退货；开发者根据快递类型传 1 或 2
}
type SubmitOrderLogisticsRequestDataItemEXTMainOrder struct {
	Express SubmitOrderLogisticsRequestDataItemEXTMainOrderExpress `json:"Express"` // 快递信息
}
type SubmitOrderLogisticsRequestDataItemEXT struct {
	MainOrder SubmitOrderLogisticsRequestDataItemEXTMainOrder `json:"MainOrder"` // 主订单信息（购买商品订单）
}
type SubmitOrderLogisticsRequestDataItem struct {
	BizAPPID   string                                 `json:"BizAPPID"`   // 小程序 AppKey
	CateID     int64                                  `json:"CateID"`     // 订单种类：1（实物）、2（虚拟物品）、5（快递服务类）、6（快递服务类无金额订单）、10（上门服务类）、11（上门服务类无金额订单）、15（酒店类）、20（票务类）、25（打车类）、26（打车类无金额订单）
	EXT        SubmitOrderLogisticsRequestDataItemEXT `json:"EXT"`        // 扩展信息
	ResourceID string                                 `json:"ResourceID"` // 开发者接入的唯一订单 ID
}

// 响应结构体

type SubmitOrderLogisticsResponsedataItem struct {
	BizAppID     string `json:"biz_app_id"`    // POST 请求参数中 BizAPPID
	CateID       string `json:"cate_id"`       // POST 请求参数中 CateID
	ResourceID   string `json:"resource_id"`   // POST 请求参数中 ResourceID
	RowsAffected string `json:"rows_affected"` // 请求受影响行数（即请求是否成功， 0 为失败，非 0 为成功）
}

type SubmitOrderLogisticsResponse struct {
	Data      []SubmitOrderLogisticsResponsedataItem `json:"data"`       // 响应对象
	Errno     int64                                  `json:"errno"`      // 错误码
	ErrMsg    string                                 `json:"msg"`        // 错误信息
	ErrorCode int64                                  `json:"error_code"` // openapi 错误码
	ErrorMsg  string                                 `json:"error_msg"`  // openapi 错误信息
}

// SubmitOrderLogistics
func SubmitOrderLogistics(params *SubmitOrderLogisticsRequest) ([]SubmitOrderLogisticsResponsedataItem, error) {
	var (
		err        error
		defaultRet []SubmitOrderLogisticsResponsedataItem
	)
	respData := &SubmitOrderLogisticsResponse{}

	client := NewHTTPClient().
		SetContentType(ContentTypeJSON).
		SetConverterType(ConverterTypeJSON).
		SetMethod("POST").
		SetScheme(SCHEME).
		SetHost(OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/ordercenter/app/add/main/logistics")
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

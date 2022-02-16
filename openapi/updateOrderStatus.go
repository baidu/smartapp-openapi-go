package openapi

import (
	"encoding/json"
	"fmt"
)

// UpdateOrderStatusRequest 请求结构体
type UpdateOrderStatusRequest struct {
	AccessToken string                             // 接口调用凭证
	OpenID      string                             // 用户 openId
	SceneID     string                             // 百度收银台分配的平台订单 ID，通知支付状态接口返回的 orderId
	SceneType   int64                              // 支付场景类型，开发者请默认传 2
	PmAppKey    string                             // 调起百度收银台的支付服务 appKey
	Data        []UpdateOrderStatusRequestDataItem // 请求数据
}

type UpdateOrderStatusRequestDataItem struct {
	BizAPPID   string `json:"BizAPPID"`   // 小程序 AppKey
	CateID     int64  `json:"CateID"`     // 订单种类：1（实物）、2（虚拟物品）、5（快递服务类）、6（快递服务类无金额订单）、10（上门服务类）、11（上门服务类无金额订单）、15（酒店类）、20（票务类）、25（打车类）、26（打车类无金额订单）
	ResourceID string `json:"ResourceID"` // 开发者接入的唯一订单 ID
	Status     int64  `json:"Status"`     // 订单状态，其值根据CateID不同有不同的定义。CateID = 1 实物订单、CateID = 2 虚拟物品订单、CateID = 5 快递服务类订单、CateID = 6 快递服务类无金额订单、CateID = 10 上门服务类订单、CateID = 11 上门服务类无金额订单、CateID = 15 酒店类订单、CateID = 20 出行票务类订单、CateID = 25 打车类订单、CateID = 26 打车类无金额订单
}

// 响应结构体

type UpdateOrderStatusResponsedataItem struct {
	BizAppID     string `json:"biz_app_id"`    // POST 请求参数中 BizAPPID
	CateID       string `json:"cate_id"`       // POST 请求参数中 CateID
	ResourceID   string `json:"resource_id"`   // POST 请求参数中 ResourceID
	RowsAffected string `json:"rows_affected"` // 请求受影响行数（即请求是否成功， 0 为失败，非 0 为成功）
}

type UpdateOrderStatusResponse struct {
	Data      []UpdateOrderStatusResponsedataItem `json:"data"`       // 响应对象
	Errno     int64                               `json:"errno"`      // 错误码
	ErrMsg    string                              `json:"msg"`        // 错误信息
	ErrorCode int64                               `json:"error_code"` // openapi 错误码
	ErrorMsg  string                              `json:"error_msg"`  // openapi 错误信息
}

// UpdateOrderStatus
func UpdateOrderStatus(params *UpdateOrderStatusRequest) ([]UpdateOrderStatusResponsedataItem, error) {
	var (
		err        error
		defaultRet []UpdateOrderStatusResponsedataItem
	)
	respData := &UpdateOrderStatusResponse{}

	client := NewHTTPClient().
		SetContentType(ContentTypeJSON).
		SetConverterType(ConverterTypeJSON).
		SetMethod("POST").
		SetScheme(SCHEME).
		SetHost(OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/ordercenter/app/update/main/status")
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

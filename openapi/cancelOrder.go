package openapi

import (
	"fmt"
)

// CancelOrderRequest 请求结构体
type CancelOrderRequest struct {
	AccessToken string // 接口调用凭证
	OrderID     int64  // 百度订单 ID
	PmAppKey    string // 调起百度收银台的支付服务 appKey
}

// 响应结构体

type CancelOrderResponse struct {
	Data      bool   `json:"data"`       // 响应对象
	Errno     int64  `json:"errno"`      // 错误码
	ErrMsg    string `json:"msg"`        // 错误信息
	ErrorCode int64  `json:"error_code"` // openapi 错误码
	ErrorMsg  string `json:"error_msg"`  // openapi 错误信息
}

// CancelOrder
func CancelOrder(params *CancelOrderRequest) (bool, error) {
	var (
		err        error
		defaultRet bool
	)
	respData := &CancelOrderResponse{}

	client := NewHTTPClient().
		SetContentType(ContentTypeForm).
		SetConverterType(ConverterTypeJSON).
		SetMethod("GET").
		SetScheme(SCHEME).
		SetHost(OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/pay/paymentservice/cancelOrder")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("orderId", fmt.Sprintf("%v", params.OrderID))
	client.AddGetParam("pmAppKey", params.PmAppKey)
	client.AddGetParam("sp_sdk_ver", SDKVERSION)
	client.AddGetParam("sp_sdk_lang", SDKLANG)

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

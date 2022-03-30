package pay

import (
	"github.com/baidu/smartapp-openapi-go/utils"
)

// CloseOrderRequest 请求结构体
type CloseOrderRequest struct {
	AccessToken string // 第三方平台的接口调用凭据
	OrderID     string // 百度订单 ID
	PmAppKey    string // 调起百度收银台的支付服务 appKey
}

// 响应结构体

type CloseOrderResponse struct {
	Data      bool   `json:"data"`       // 响应参数
	Errno     int64  `json:"errno"`      // 状态码
	ErrMsg    string `json:"msg"`        // 错误信息
	ErrorCode int64  `json:"error_code"` // openapi 错误码
	ErrorMsg  string `json:"error_msg"`  // openapi 错误信息
}

// CloseOrder
func CloseOrder(params *CloseOrderRequest) (bool, error) {
	var (
		err        error
		defaultRet bool
	)
	respData := &CloseOrderResponse{}

	client := utils.NewHTTPClient().
		SetContentType(utils.ContentTypeForm).
		SetConverterType(utils.ConverterTypeJSON).
		SetMethod("GET").
		SetScheme(utils.SCHEME).
		SetHost(utils.OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/pay/paymentservice/tp/cancelOrder")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("orderId", params.OrderID)
	client.AddGetParam("pmAppKey", params.PmAppKey)
	client.AddGetParam("sp_sdk_ver", utils.SDKVERSION)
	client.AddGetParam("sp_sdk_lang", utils.SDKLANG)

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

package pay

import (
	"fmt"

	"github.com/baidu/smartapp-openapi-go/utils"
)

// GetOrderRefundRequest 请求结构体
type GetOrderRefundRequest struct {
	AccessToken string // 第三方平台的接口调用凭据
	TpOrderID   string // 开发者订单 ID
	UserID      int64  // 百度收银台用户 ID
	PmAppKey    string // 调起百度收银台的支付服务 appKey
}

// 响应结构体

type GetOrderRefundResponsedataItem struct {
	BizRefundBatchID string `json:"bizRefundBatchId"` // 开发者退款批次id
	OrderID          int64  `json:"orderId"`          // 退款订单号
	RefundBatchID    int64  `json:"refundBatchId"`    // 退款批次id
	RefundStatus     int64  `json:"refundStatus"`     // 退款状态 1 退款中 2 退款成功 3 退款失败
	UserID           int64  `json:"userId"`           // 退款用户id
}

type GetOrderRefundResponse struct {
	Data      []GetOrderRefundResponsedataItem `json:"data"`       // 响应参数
	Errno     int64                            `json:"errno"`      // 状态码
	ErrMsg    string                           `json:"msg"`        // 错误信息
	ErrorCode int64                            `json:"error_code"` // openapi 错误码
	ErrorMsg  string                           `json:"error_msg"`  // openapi 错误信息
}

// GetOrderRefund
func GetOrderRefund(params *GetOrderRefundRequest) ([]GetOrderRefundResponsedataItem, error) {
	var (
		err        error
		defaultRet []GetOrderRefundResponsedataItem
	)
	respData := &GetOrderRefundResponse{}

	client := utils.NewHTTPClient().
		SetContentType(utils.ContentTypeForm).
		SetConverterType(utils.ConverterTypeJSON).
		SetMethod("GET").
		SetScheme(utils.SCHEME).
		SetHost(utils.OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/pay/paymentservice/tp/findOrderRefund")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("tpOrderId", params.TpOrderID)
	client.AddGetParam("userId", fmt.Sprintf("%v", params.UserID))
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

package openapi

import (
	"fmt"

)

// FindOrderRefundRequest 请求结构体
type FindOrderRefundRequest struct {
	AccessToken string // 接口调用凭证
	TpOrderID   string // 开发者订单 ID
	UserID      int64  // 百度收银台用户 ID
	PmAppKey    string // 调起百度收银台的支付服务 appKey
}

// 响应结构体

type FindOrderRefundResponsedataItem struct {
	BizRefundBatchID string `json:"bizRefundBatchId"` // 开发者退款批次id
	OrderID          int64  `json:"orderId"`          // 退款订单号
	RefundBatchID    int64  `json:"refundBatchId"`    // 退款批次id
	RefundStatus     int64  `json:"refundStatus"`     // 退款状态 1 退款中 2 退款成功 3 退款失败
	UserID           int64  `json:"userId"`           // 退款用户id
}

type FindOrderRefundResponse struct {
	Data      []FindOrderRefundResponsedataItem `json:"data"`       // 响应对象
	Errno     int64                             `json:"errno"`      // 错误码
	ErrMsg    string                            `json:"msg"`        // 错误信息
	ErrorCode int64                             `json:"error_code"` // openapi 错误码
	ErrorMsg  string                            `json:"error_msg"`  // openapi 错误信息
}

// FindOrderRefund
func FindOrderRefund(params *FindOrderRefundRequest) ([]FindOrderRefundResponsedataItem, error) {
	var (
		err        error
		defaultRet []FindOrderRefundResponsedataItem
	)
	respData := &FindOrderRefundResponse{}

	client := NewHTTPClient().
		SetContentType(ContentTypeForm).
		SetConverterType(ConverterTypeJSON).
		SetMethod("GET").
		SetScheme(SCHEME).
		SetHost(OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/pay/paymentservice/findOrderRefund")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("tpOrderId", params.TpOrderID)
	client.AddGetParam("userId", fmt.Sprintf("%v", params.UserID))
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

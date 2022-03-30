package pay

import (
	"fmt"

	"github.com/baidu/smartapp-openapi-go/utils"
)

// ApplyOrderRefundRequest 请求结构体
type ApplyOrderRefundRequest struct {
	AccessToken      string      // 第三方平台的接口调用凭据
	ApplyRefundMoney interface{} // 退款金额（单位：分），该字段最大不能超过支付回调中的总金额（totalMoney）1.如不填金额时，默认整单发起退款2.含有百度平台营销的订单，目前只支持整单发起退款，不支持部分多次退款
	BizRefundBatchID string      // 开发者退款批次
	IsSkipAudit      int64       // 是否跳过审核，不需要百度请求开发者退款审核请传 1，默认为0； 0：不跳过开发者业务方审核；1：跳过开发者业务方审核。若不跳过审核，请对接请求业务方退款审核接口
	OrderID          int64       // 百度收银台订单 ID
	RefundReason     string      // 退款原因
	RefundType       int64       // 退款类型 1：用户发起退款；2：开发者业务方客服退款；3：开发者服务异常退款。
	TpOrderID        string      // 开发者订单 ID
	UserID           int64       // 百度收银台用户 ID
	RefundNotifyURL  interface{} // 退款通知 url ，不传时默认为在开发者后台配置的 url
	PmAppKey         string      // 调起百度收银台的支付服务 appKey
}

// 响应结构体

type ApplyOrderRefundResponsedata struct {
	RefundBatchID  string `json:"refundBatchId"`  // 平台退款批次号
	RefundPayMoney int64  `json:"refundPayMoney"` // 平台可退退款金额【分为单位】
}

type ApplyOrderRefundResponse struct {
	Data      ApplyOrderRefundResponsedata `json:"data"`       // 响应参数
	Errno     int64                        `json:"errno"`      // 状态码
	ErrMsg    string                       `json:"msg"`        // 错误信息
	ErrorCode int64                        `json:"error_code"` // openapi 错误码
	ErrorMsg  string                       `json:"error_msg"`  // openapi 错误信息
}

// ApplyOrderRefund
func ApplyOrderRefund(params *ApplyOrderRefundRequest) (*ApplyOrderRefundResponsedata, error) {
	var (
		err        error
		defaultRet *ApplyOrderRefundResponsedata
	)
	respData := &ApplyOrderRefundResponse{}

	client := utils.NewHTTPClient().
		SetContentType(utils.ContentTypeForm).
		SetConverterType(utils.ConverterTypeJSON).
		SetMethod("POST").
		SetScheme(utils.SCHEME).
		SetHost(utils.OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/pay/paymentservice/tp/applyOrderRefund")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("sp_sdk_ver", utils.SDKVERSION)
	client.AddGetParam("sp_sdk_lang", utils.SDKLANG)
	client.AddPostParam("applyRefundMoney", params.ApplyRefundMoney)
	client.AddPostParam("bizRefundBatchId", params.BizRefundBatchID)
	client.AddPostParam("isSkipAudit", fmt.Sprintf("%v", params.IsSkipAudit))
	client.AddPostParam("orderId", fmt.Sprintf("%v", params.OrderID))
	client.AddPostParam("refundReason", params.RefundReason)
	client.AddPostParam("refundType", fmt.Sprintf("%v", params.RefundType))
	client.AddPostParam("tpOrderId", params.TpOrderID)
	client.AddPostParam("userId", fmt.Sprintf("%v", params.UserID))
	client.AddPostParam("refundNotifyUrl", params.RefundNotifyURL)
	client.AddPostParam("pmAppKey", params.PmAppKey)

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
	return &respData.Data, nil
}

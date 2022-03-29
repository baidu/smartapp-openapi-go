package pay

import (
	"github.com/baidu/smartapp-openapi-go/utils"
)

// SubmitDevelopConfigRequest 请求结构体
type SubmitDevelopConfigRequest struct {
	AccessToken    string // 授权小程序的接口调用凭据
	TpPublicKey    string // 开发者公钥。参见RSA公私钥生成
	PayNotifyURL   string // 支付回调地址。
	RefundAuditURL string // 退款审核地址。
	RefundSuccURL  string // 退款回调地址。
}

// 响应结构体

type SubmitDevelopConfigResponsedata struct {
}

type SubmitDevelopConfigResponse struct {
	Data      SubmitDevelopConfigResponsedata `json:"data"`       // 响应参数
	Errno     int64                           `json:"errno"`      // 状态码
	ErrMsg    string                          `json:"msg"`        // 错误信息
	ErrorCode int64                           `json:"error_code"` // openapi 错误码
	ErrorMsg  string                          `json:"error_msg"`  // openapi 错误信息
}

// SubmitDevelopConfig
func SubmitDevelopConfig(params *SubmitDevelopConfigRequest) (*SubmitDevelopConfigResponsedata, error) {
	var (
		err        error
		defaultRet *SubmitDevelopConfigResponsedata
	)
	respData := &SubmitDevelopConfigResponse{}

	client := utils.NewHTTPClient().
		SetContentType(utils.ContentTypeForm).
		SetConverterType(utils.ConverterTypeJSON).
		SetMethod("POST").
		SetScheme(utils.SCHEME).
		SetHost(utils.OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/pay/developconfig/submit")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("sp_sdk_ver", utils.SDKVERSION)
	client.AddGetParam("sp_sdk_lang", utils.SDKLANG)
	client.AddPostParam("tp_public_key", params.TpPublicKey)
	client.AddPostParam("pay_notify_url", params.PayNotifyURL)
	client.AddPostParam("refund_audit_url", params.RefundAuditURL)
	client.AddPostParam("refund_succ_url", params.RefundSuccURL)

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

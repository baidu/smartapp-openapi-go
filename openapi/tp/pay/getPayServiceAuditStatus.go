package pay

import (
	"github.com/baidu/smartapp-openapi-go/utils"
)

// GetPayServiceAuditStatusRequest 请求结构体
type GetPayServiceAuditStatusRequest struct {
	AccessToken string // 授权小程序的接口调用凭据
}

// 响应结构体

type GetPayServiceAuditStatusResponsedata struct {
	FailReason string `json:"fail_reason"` // 失败原因。
	Status     int64  `json:"status"`      // 开通状态 0：新建 1：审核中 2：审核通过 3：驳回
}

type GetPayServiceAuditStatusResponse struct {
	Data      GetPayServiceAuditStatusResponsedata `json:"data"`       // 响应参数
	Errno     int64                                `json:"errno"`      // 状态码
	ErrMsg    string                               `json:"msg"`        // 错误信息
	ErrorCode int64                                `json:"error_code"` // openapi 错误码
	ErrorMsg  string                               `json:"error_msg"`  // openapi 错误信息
}

// GetPayServiceAuditStatus
func GetPayServiceAuditStatus(params *GetPayServiceAuditStatusRequest) (*GetPayServiceAuditStatusResponsedata, error) {
	var (
		err        error
		defaultRet *GetPayServiceAuditStatusResponsedata
	)
	respData := &GetPayServiceAuditStatusResponse{}

	client := utils.NewHTTPClient().
		SetContentType(utils.ContentTypeForm).
		SetConverterType(utils.ConverterTypeJSON).
		SetMethod("GET").
		SetScheme(utils.SCHEME).
		SetHost(utils.OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/pay/paymentservice/auditstatus")
	client.AddGetParam("access_token", params.AccessToken)
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
	return &respData.Data, nil
}

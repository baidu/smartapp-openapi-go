package pay

import (
	"github.com/baidu/smartapp-openapi-go/utils"
)

// GetPayAccountAuditStatusRequest 请求结构体
type GetPayAccountAuditStatusRequest struct {
	AccessToken string // 授权小程序的接口调用凭据
}

// 响应结构体

type GetPayAccountAuditStatusResponsedata struct {
	Status        int64  `json:"status"`         // 用户提交的审核信息：0:初始状态；1:编辑中 2:提审 3:未通过 4:审核通过 5:回填过协议id
	VerifyMessage string `json:"verify_message"` // 失败原因。
}

type GetPayAccountAuditStatusResponse struct {
	Data      GetPayAccountAuditStatusResponsedata `json:"data"`       // 响应参数
	Errno     int64                                `json:"errno"`      // 状态码
	ErrMsg    string                               `json:"msg"`        // 错误信息
	ErrorCode int64                                `json:"error_code"` // openapi 错误码
	ErrorMsg  string                               `json:"error_msg"`  // openapi 错误信息
}

// GetPayAccountAuditStatus
func GetPayAccountAuditStatus(params *GetPayAccountAuditStatusRequest) (*GetPayAccountAuditStatusResponsedata, error) {
	var (
		err        error
		defaultRet *GetPayAccountAuditStatusResponsedata
	)
	respData := &GetPayAccountAuditStatusResponse{}

	client := utils.NewHTTPClient().
		SetContentType(utils.ContentTypeForm).
		SetConverterType(utils.ConverterTypeJSON).
		SetMethod("GET").
		SetScheme(utils.SCHEME).
		SetHost(utils.OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/pay/account/auditstatus")
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

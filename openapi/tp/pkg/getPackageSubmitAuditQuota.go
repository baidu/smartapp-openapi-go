package pkg

import (
	"github.com/baidu/smartapp-openapi-go/utils"
)

// GetPackageSubmitAuditQuotaRequest 请求结构体
type GetPackageSubmitAuditQuotaRequest struct {
	AccessToken string // 授权小程序的接口调用凭据
}

// 响应结构体

type GetPackageSubmitAuditQuotaResponsedata struct {
	QuotaTimes  int64 `json:"quota_times"`  // 单个周期内额度上限，-1：无限制
	QuotaType   int64 `json:"quota_type"`   // 额度限制，周期类型：2：周
	RemainTimes int64 `json:"remain_times"` // 当前周期内剩余额度，-1：无限制
}

type GetPackageSubmitAuditQuotaResponse struct {
	Data      GetPackageSubmitAuditQuotaResponsedata `json:"data"`       // 响应参数
	Errno     int64                                  `json:"errno"`      // 状态码
	ErrMsg    string                                 `json:"msg"`        // 错误信息
	ErrorCode int64                                  `json:"error_code"` // openapi 错误码
	ErrorMsg  string                                 `json:"error_msg"`  // openapi 错误信息
}

// GetPackageSubmitAuditQuota
func GetPackageSubmitAuditQuota(params *GetPackageSubmitAuditQuotaRequest) (*GetPackageSubmitAuditQuotaResponsedata, error) {
	var (
		err        error
		defaultRet *GetPackageSubmitAuditQuotaResponsedata
	)
	respData := &GetPackageSubmitAuditQuotaResponse{}

	client := utils.NewHTTPClient().
		SetContentType(utils.ContentTypeForm).
		SetConverterType(utils.ConverterTypeJSON).
		SetMethod("GET").
		SetScheme(utils.SCHEME).
		SetHost(utils.OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/package/auditquota")
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

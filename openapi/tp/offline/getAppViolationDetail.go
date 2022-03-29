package offline

import (
	"github.com/baidu/smartapp-openapi-go/utils"
)

// GetAppViolationDetailRequest 请求结构体
type GetAppViolationDetailRequest struct {
	AccessToken string // 授权小程序的接口调用凭据
}

// 响应结构体

type GetAppViolationDetailResponsedataproblemDetailsItem struct {
	AuditReason string `json:"auditReason"` // 审核原因
	AuditShot   string `json:"auditShot"`   // 审核截图
	ProblemDesc string `json:"problemDesc"` // 问题描述
	ProblemPath string `json:"problemPath"` // 问题 path
	ScreenShot  string `json:"screenShot"`  // 问题截图
	Status      int64  `json:"status"`      // 问题状态
	TaskID      int64  `json:"taskId"`      // 任务 ID
}

type GetAppViolationDetailResponsedata struct {
	AppID          int64                                                 `json:"appId"`          // 小程序 ID
	ProblemDetails []GetAppViolationDetailResponsedataproblemDetailsItem `json:"problemDetails"` //
}

type GetAppViolationDetailResponse struct {
	Data      GetAppViolationDetailResponsedata `json:"data"`       // 响应参数
	Errno     int64                             `json:"errno"`      // 状态码
	ErrMsg    string                            `json:"msg"`        // 错误信息
	ErrorCode int64                             `json:"error_code"` // openapi 错误码
	ErrorMsg  string                            `json:"error_msg"`  // openapi 错误信息
}

// GetAppViolationDetail
func GetAppViolationDetail(params *GetAppViolationDetailRequest) (*GetAppViolationDetailResponsedata, error) {
	var (
		err        error
		defaultRet *GetAppViolationDetailResponsedata
	)
	respData := &GetAppViolationDetailResponse{}

	client := utils.NewHTTPClient().
		SetContentType(utils.ContentTypeForm).
		SetConverterType(utils.ConverterTypeJSON).
		SetMethod("GET").
		SetScheme(utils.SCHEME).
		SetHost(utils.OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/app/violation")
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

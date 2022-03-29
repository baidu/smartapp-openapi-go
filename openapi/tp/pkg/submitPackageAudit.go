package pkg

import (
	"github.com/baidu/smartapp-openapi-go/utils"
)

// SubmitPackageAuditRequest 请求结构体
type SubmitPackageAuditRequest struct {
	AccessToken  string      // 授权小程序的接口调用凭据
	Content      string      // 送审描述
	PackageID    string      // 包 id ，获取方式请参考获取小程序包列表
	Remark       string      // 备注
	TestAccount  interface{} // 可以向审核人员提供的测试帐号
	TestPassword interface{} // 测试帐号对应的密码
}

// 响应结构体

type SubmitPackageAuditResponsedata struct {
}

type SubmitPackageAuditResponse struct {
	Data      SubmitPackageAuditResponsedata `json:"data"`       // 响应参数
	Errno     int64                          `json:"errno"`      // 状态码
	ErrMsg    string                         `json:"msg"`        // 错误信息
	ErrorCode int64                          `json:"error_code"` // openapi 错误码
	ErrorMsg  string                         `json:"error_msg"`  // openapi 错误信息
}

// SubmitPackageAudit
func SubmitPackageAudit(params *SubmitPackageAuditRequest) (*SubmitPackageAuditResponsedata, error) {
	var (
		err        error
		defaultRet *SubmitPackageAuditResponsedata
	)
	respData := &SubmitPackageAuditResponse{}

	client := utils.NewHTTPClient().
		SetContentType(utils.ContentTypeForm).
		SetConverterType(utils.ConverterTypeJSON).
		SetMethod("POST").
		SetScheme(utils.SCHEME).
		SetHost(utils.OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/package/submitaudit")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("sp_sdk_ver", utils.SDKVERSION)
	client.AddGetParam("sp_sdk_lang", utils.SDKLANG)
	client.AddPostParam("content", params.Content)
	client.AddPostParam("package_id", params.PackageID)
	client.AddPostParam("remark", params.Remark)
	client.AddPostParam("test_account", params.TestAccount)
	client.AddPostParam("test_password", params.TestPassword)

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

package pkg

import (
	"github.com/baidu/smartapp-openapi-go/utils"
)

// UploadPackageRequest 请求结构体
type UploadPackageRequest struct {
	AccessToken  string      // 授权小程序的接口调用凭据
	TemplateID   string      // 代码库中的代码模板 id ，可以在第三方平台-模板管理-模板库 查看到模板 id
	ExtJSON      string      // 第三方自定义的配置
	UserDesc     string      // 代码描述，开发者可自定义。
	UserVersion  string      // 代码版本号，开发者可自定义。
	TestAccount  interface{} // 设置直接送审( ext_json 中的 directCommit 字段为 true 时)，可以向审核人员提供的测试帐号。
	TestPassword interface{} // 测试帐号对应的密码。
}

// 响应结构体

type UploadPackageResponsedata struct {
}

type UploadPackageResponse struct {
	Data      UploadPackageResponsedata `json:"data"`       // 响应参数
	Errno     int64                     `json:"errno"`      // 状态码
	ErrMsg    string                    `json:"msg"`        // 错误信息
	ErrorCode int64                     `json:"error_code"` // openapi 错误码
	ErrorMsg  string                    `json:"error_msg"`  // openapi 错误信息
}

// UploadPackage
func UploadPackage(params *UploadPackageRequest) (*UploadPackageResponsedata, error) {
	var (
		err        error
		defaultRet *UploadPackageResponsedata
	)
	respData := &UploadPackageResponse{}

	client := utils.NewHTTPClient().
		SetContentType(utils.ContentTypeForm).
		SetConverterType(utils.ConverterTypeJSON).
		SetMethod("POST").
		SetScheme(utils.SCHEME).
		SetHost(utils.OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/package/upload")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("sp_sdk_ver", utils.SDKVERSION)
	client.AddGetParam("sp_sdk_lang", utils.SDKLANG)
	client.AddPostParam("template_id", params.TemplateID)
	client.AddPostParam("ext_json", params.ExtJSON)
	client.AddPostParam("user_desc", params.UserDesc)
	client.AddPostParam("user_version", params.UserVersion)
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

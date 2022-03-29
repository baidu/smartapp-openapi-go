package app

import (
	"github.com/baidu/smartapp-openapi-go/utils"
)

// CheckAppNickNameRequest 请求结构体
type CheckAppNickNameRequest struct {
	AccessToken   string      // 授权小程序的接口调用凭据
	AppName       string      // 小程序名字
	QualMaterials interface{} // 如果小程序名称包含品牌词，需要上传品牌资质证明。
}

// 响应结构体

type CheckAppNickNameResponsedata struct {
	CheckResult    int64    `json:"checkResult"`    // 检测结果码
	CheckWords     []string `json:"checkWords"`     // 命中关键词
	OptionalFields []int64  `json:"optionalFields"` // 当前名称可上传的资质字段（qual_materials参数）
	RequiredFields []int64  `json:"requiredFields"` // 当前名称必须要上传的资质字段（qual_materials参数）
}

type CheckAppNickNameResponse struct {
	Data      CheckAppNickNameResponsedata `json:"data"`       // 响应参数
	Errno     int64                        `json:"errno"`      // 状态码
	ErrMsg    string                       `json:"msg"`        // 错误信息
	ErrorCode int64                        `json:"error_code"` // openapi 错误码
	ErrorMsg  string                       `json:"error_msg"`  // openapi 错误信息
}

// CheckAppNickName
func CheckAppNickName(params *CheckAppNickNameRequest) (*CheckAppNickNameResponsedata, error) {
	var (
		err        error
		defaultRet *CheckAppNickNameResponsedata
	)
	respData := &CheckAppNickNameResponse{}

	client := utils.NewHTTPClient().
		SetContentType(utils.ContentTypeForm).
		SetConverterType(utils.ConverterTypeJSON).
		SetMethod("GET").
		SetScheme(utils.SCHEME).
		SetHost(utils.OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/app/checknamewithqual")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("app_name", params.AppName)
	client.AddGetParam("qual_materials", params.QualMaterials)
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

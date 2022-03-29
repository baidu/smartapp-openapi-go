package app

import (
	"github.com/baidu/smartapp-openapi-go/utils"
)

// SetAppNickNameRequest 请求结构体
type SetAppNickNameRequest struct {
	AccessToken   string      // 授权小程序的接口调用凭据
	NickName      string      // 小程序名字
	QualMaterials interface{} // 如果小程序名称包含品牌词，需要上传品牌资质证明。要求图片链接来自于图片上传接口返回的 url。
}

// 响应结构体

type SetAppNickNameResponsedataItem struct {
	ExamineID int64 `json:"examine_id"` //
}

type SetAppNickNameResponse struct {
	Data      []SetAppNickNameResponsedataItem `json:"data"`       // 响应参数
	Errno     int64                            `json:"errno"`      // 状态码
	ErrMsg    string                           `json:"msg"`        // 错误信息
	ErrorCode int64                            `json:"error_code"` // openapi 错误码
	ErrorMsg  string                           `json:"error_msg"`  // openapi 错误信息
}

// SetAppNickName
func SetAppNickName(params *SetAppNickNameRequest) ([]SetAppNickNameResponsedataItem, error) {
	var (
		err        error
		defaultRet []SetAppNickNameResponsedataItem
	)
	respData := &SetAppNickNameResponse{}

	client := utils.NewHTTPClient().
		SetContentType(utils.ContentTypeForm).
		SetConverterType(utils.ConverterTypeJSON).
		SetMethod("POST").
		SetScheme(utils.SCHEME).
		SetHost(utils.OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/app/setnicknamewithqual")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("sp_sdk_ver", utils.SDKVERSION)
	client.AddGetParam("sp_sdk_lang", utils.SDKLANG)
	client.AddPostParam("nick_name", params.NickName)
	client.AddPostParam("qual_materials", params.QualMaterials)

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
	return respData.Data, nil
}

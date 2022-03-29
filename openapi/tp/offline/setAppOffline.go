package offline

import (
	"github.com/baidu/smartapp-openapi-go/utils"
)

// SetAppOfflineRequest 请求结构体
type SetAppOfflineRequest struct {
	AccessToken     string      // 授权小程序的接口调用凭据
	AppName         interface{} // 小程序名称
	AppDesc         interface{} // 小程序描述
	PhotoAddr       interface{} // 小程序图片
	AppNameMaterial interface{} // 名称相关物料
}

// 响应结构体

type SetAppOfflineResponsedata struct {
	ExamineID int64 `json:"examine_id"` // 审核ID
}

type SetAppOfflineResponse struct {
	Data      SetAppOfflineResponsedata `json:"data"`       // 响应参数
	Errno     int64                     `json:"errno"`      // 状态码
	ErrMsg    string                    `json:"msg"`        // 错误信息
	ErrorCode int64                     `json:"error_code"` // openapi 错误码
	ErrorMsg  string                    `json:"error_msg"`  // openapi 错误信息
}

// SetAppOffline
func SetAppOffline(params *SetAppOfflineRequest) (*SetAppOfflineResponsedata, error) {
	var (
		err        error
		defaultRet *SetAppOfflineResponsedata
	)
	respData := &SetAppOfflineResponse{}

	client := utils.NewHTTPClient().
		SetContentType(utils.ContentTypeForm).
		SetConverterType(utils.ConverterTypeJSON).
		SetMethod("POST").
		SetScheme(utils.SCHEME).
		SetHost(utils.OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/app/offline/update")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("sp_sdk_ver", utils.SDKVERSION)
	client.AddGetParam("sp_sdk_lang", utils.SDKLANG)
	client.AddPostParam("app_name", params.AppName)
	client.AddPostParam("app_desc", params.AppDesc)
	client.AddPostParam("photo_addr", params.PhotoAddr)
	client.AddPostParam("app_name_material", params.AppNameMaterial)

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

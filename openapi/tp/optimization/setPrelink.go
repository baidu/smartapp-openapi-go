package optimization

import (
	"github.com/baidu/smartapp-openapi-go/utils"
)

// SetPrelinkRequest 请求结构体
type SetPrelinkRequest struct {
	AccessToken string // 授权小程序的接口调用凭据
	PrelinkURL  string // prelink 链接
}

// 响应结构体

type SetPrelinkResponsedatacheckResulthttpsBaikeBaiduCom struct {
	FailReason []string `json:"failReason"` // 失败原因
	Status     int64    `json:"status"`     // 状态 1-失败 0-成功
}

type SetPrelinkResponsedatacheckResulthttpsWwwBaiduCom struct {
	FailReason []string `json:"failReason"` // 失败原因
	Status     int64    `json:"status"`     // 状态 1-失败 0-成功
}

type SetPrelinkResponsedatacheckResultflexKey struct {
	FailReason []string `json:"failReason"` //
	Status     string   `json:"status"`     //
}

type SetPrelinkResponsedatacheckResult struct {
	FlexKey            SetPrelinkResponsedatacheckResultflexKey            `json:"flexKey"`                 //
	HTTPSBaikeBaiduCom SetPrelinkResponsedatacheckResulthttpsBaikeBaiduCom `json:"https://baike.baidu.com"` //
	HTTPSWwwBaiduCom   SetPrelinkResponsedatacheckResulthttpsWwwBaiduCom   `json:"https://www.baidu.com"`   //
}

type SetPrelinkResponsedata struct {
	CheckResult SetPrelinkResponsedatacheckResult `json:"checkResult"` // 域名检查返回
	Status      int64                             `json:"status"`      // 状态 1-失败 0-成功
}

type SetPrelinkResponse struct {
	Data      SetPrelinkResponsedata `json:"data"`       // 响应参数
	Errno     int64                  `json:"errno"`      // 状态码
	ErrMsg    string                 `json:"msg"`        // 错误信息
	ErrorCode int64                  `json:"error_code"` // openapi 错误码
	ErrorMsg  string                 `json:"error_msg"`  // openapi 错误信息
}

// SetPrelink
func SetPrelink(params *SetPrelinkRequest) (*SetPrelinkResponsedata, error) {
	var (
		err        error
		defaultRet *SetPrelinkResponsedata
	)
	respData := &SetPrelinkResponse{}

	client := utils.NewHTTPClient().
		SetContentType(utils.ContentTypeForm).
		SetConverterType(utils.ConverterTypeJSON).
		SetMethod("POST").
		SetScheme(utils.SCHEME).
		SetHost(utils.OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/prelink/set")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("sp_sdk_ver", utils.SDKVERSION)
	client.AddGetParam("sp_sdk_lang", utils.SDKLANG)
	client.AddPostParam("prelink_url", params.PrelinkURL)

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

package domain

import (
	"github.com/baidu/smartapp-openapi-go/utils"
)

// ModifyBusinessDomainRequest 请求结构体
type ModifyBusinessDomainRequest struct {
	AccessToken   string      // 第三方平台的接口调用凭据
	Action        interface{} // add添加, delete删除, set覆盖, get获取。当参数是get时不需要填四个域名字段，如果没有action字段参数，则默认将开放平台第三方登记的小程序业务域名全部添加到授权的小程序中。
	WebViewDomain interface{} // 小程序业务域名，多个时用,分割，当action参数是get时不需要此字段
}

// 响应结构体

type ModifyBusinessDomainResponse struct {
	Data      []string `json:"data"`       // 响应参数
	Errno     int64    `json:"errno"`      // 状态码
	ErrMsg    string   `json:"msg"`        // 错误信息
	ErrorCode int64    `json:"error_code"` // openapi 错误码
	ErrorMsg  string   `json:"error_msg"`  // openapi 错误信息
}

// ModifyBusinessDomain
func ModifyBusinessDomain(params *ModifyBusinessDomainRequest) ([]string, error) {
	var (
		err        error
		defaultRet []string
	)
	respData := &ModifyBusinessDomainResponse{}

	client := utils.NewHTTPClient().
		SetContentType(utils.ContentTypeForm).
		SetConverterType(utils.ConverterTypeJSON).
		SetMethod("POST").
		SetScheme(utils.SCHEME).
		SetHost(utils.OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/tp/modifywebviewdomain")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("sp_sdk_ver", utils.SDKVERSION)
	client.AddGetParam("sp_sdk_lang", utils.SDKLANG)
	client.AddPostParam("action", params.Action)
	client.AddPostParam("web_view_domain", params.WebViewDomain)

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

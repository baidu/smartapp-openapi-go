package app

import (
	"github.com/baidu/smartapp-openapi-go/utils"
)

// SetAppDomainRequest 请求结构体
type SetAppDomainRequest struct {
	AccessToken    string      // 授权小程序的接口调用凭据
	Action         interface{} // add添加, delete删除, set覆盖, get获取。当参数是get时不需要填四个域名字段，如果没有action字段参数，则默认将开放平台第三方登记的小程序业务域名全部添加到授权的小程序中。
	DownloadDomain interface{} // download合法域名，多个时用,分割，当action参数是get时不需要此字段
	RequestDomain  interface{} // request合法域名，多个时用,分割，当action参数是get时不需要此字段。
	SocketDomain   interface{} // socket合法域名，多个时用,分割，当action参数是get时不需要此字段。
	UploadDomain   interface{} // upload合法域名，多个时用,分割，当action参数是get时不需要此字段。
}

// 响应结构体

type SetAppDomainResponsedata struct {
	DownloadDomain string `json:"download_domain"` // download合法域名
	RequestDomain  string `json:"request_domain"`  // request合法域名
	SocketDomain   string `json:"socket_domain"`   // socket合法域名
	UploadDomain   string `json:"upload_domain"`   // upload合法域名
}

type SetAppDomainResponse struct {
	Data      SetAppDomainResponsedata `json:"data"`       // 响应参数
	Errno     int64                    `json:"errno"`      // 状态码
	ErrMsg    string                   `json:"msg"`        // 错误信息
	ErrorCode int64                    `json:"error_code"` // openapi 错误码
	ErrorMsg  string                   `json:"error_msg"`  // openapi 错误信息
}

// SetAppDomain
func SetAppDomain(params *SetAppDomainRequest) (*SetAppDomainResponsedata, error) {
	var (
		err        error
		defaultRet *SetAppDomainResponsedata
	)
	respData := &SetAppDomainResponse{}

	client := utils.NewHTTPClient().
		SetContentType(utils.ContentTypeForm).
		SetConverterType(utils.ConverterTypeJSON).
		SetMethod("POST").
		SetScheme(utils.SCHEME).
		SetHost(utils.OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/app/modifydomain")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("sp_sdk_ver", utils.SDKVERSION)
	client.AddGetParam("sp_sdk_lang", utils.SDKLANG)
	client.AddPostParam("action", params.Action)
	client.AddPostParam("download_domain", params.DownloadDomain)
	client.AddPostParam("request_domain", params.RequestDomain)
	client.AddPostParam("socket_domain", params.SocketDomain)
	client.AddPostParam("upload_domain", params.UploadDomain)

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

package domain

import (
	"github.com/baidu/smartapp-openapi-go/utils"
)

// DownloadDomainCheckFileRequest 请求结构体
type DownloadDomainCheckFileRequest struct {
	AccessToken string // 第三方平台的接口调用凭据
}

// 响应结构体

type DownloadDomainCheckFileResponse struct {
	ContentType string
	Data        DownloadDomainCheckFileResponsedata
	ErrorCode   int64  `json:"error_code"` // openapi 错误码
	ErrorMsg    string `json:"error_msg"`  // openapi 错误信息
}

type DownloadDomainCheckFileResponsedata []byte

// DownloadDomainCheckFile
func DownloadDomainCheckFile(params *DownloadDomainCheckFileRequest) (*DownloadDomainCheckFileResponse, error) {
	var (
		err        error
		defaultRet *DownloadDomainCheckFileResponse
	)
	respData := &DownloadDomainCheckFileResponse{}

	client := utils.NewHTTPClient().
		SetContentType(utils.ContentTypeForm).
		SetConverterType(utils.ConverterTypeStream).
		SetMethod("GET").
		SetScheme(utils.SCHEME).
		SetHost(utils.OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/tp/download/domaincer")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("sp_sdk_ver", utils.SDKVERSION)
	client.AddGetParam("sp_sdk_lang", utils.SDKLANG)

	err = client.Do()
	if err != nil {
		return defaultRet, err
	}
	err = client.Convert(&respData.Data)
	if err != nil {
		return defaultRet, err
	}
	if respData.ErrorCode != 0 {
		return defaultRet, &utils.OpenAPIError{respData.ErrorCode, respData.ErrorMsg, respData}
	}

	respData.ContentType = utils.ConverterTypeStream
	return respData, nil
}

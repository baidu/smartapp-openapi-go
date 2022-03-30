package search

import (
	"github.com/baidu/smartapp-openapi-go/utils"
)

// DeleteSubChainRequest 请求结构体
type DeleteSubChainRequest struct {
	AccessToken string // 授权小程序的接口调用凭据
	SubchainID  string // 子链 Id
}

// 响应结构体

type DeleteSubChainResponse struct {
	Data      bool   `json:"data"`       // 是否成功删除
	Errno     int64  `json:"errno"`      // 状态码
	ErrMsg    string `json:"msg"`        // 错误信息
	ErrorCode int64  `json:"error_code"` // openapi 错误码
	ErrorMsg  string `json:"error_msg"`  // openapi 错误信息
}

// DeleteSubChain
func DeleteSubChain(params *DeleteSubChainRequest) (bool, error) {
	var (
		err        error
		defaultRet bool
	)
	respData := &DeleteSubChainResponse{}

	client := utils.NewHTTPClient().
		SetContentType(utils.ContentTypeForm).
		SetConverterType(utils.ConverterTypeJSON).
		SetMethod("POST").
		SetScheme(utils.SCHEME).
		SetHost(utils.OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/subchain/delete")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("sp_sdk_ver", utils.SDKVERSION)
	client.AddGetParam("sp_sdk_lang", utils.SDKLANG)
	client.AddPostParam("subchain_id", params.SubchainID)

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

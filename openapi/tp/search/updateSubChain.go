package search

import (
	"github.com/baidu/smartapp-openapi-go/utils"
)

// UpdateSubChainRequest 请求结构体
type UpdateSubChainRequest struct {
	AccessToken string      // 授权小程序的接口调用凭据
	SubchainID  string      // 子链 Id
	ChainName   string      // 4-10个字符，说明子链的功能
	ChainDesc   string      // 8-16个字符，辅助描述子链的功能
	ChainPath   interface{} // 以“/”开头的子链对应的path路径
	Telephone   interface{} // SA类型的客服电话子链
}

// 响应结构体

type UpdateSubChainResponse struct {
	Data      bool   `json:"data"`       // 是否成功更新
	Errno     int64  `json:"errno"`      // 状态码
	ErrMsg    string `json:"msg"`        // 错误信息
	ErrorCode int64  `json:"error_code"` // openapi 错误码
	ErrorMsg  string `json:"error_msg"`  // openapi 错误信息
}

// UpdateSubChain
func UpdateSubChain(params *UpdateSubChainRequest) (bool, error) {
	var (
		err        error
		defaultRet bool
	)
	respData := &UpdateSubChainResponse{}

	client := utils.NewHTTPClient().
		SetContentType(utils.ContentTypeForm).
		SetConverterType(utils.ConverterTypeJSON).
		SetMethod("POST").
		SetScheme(utils.SCHEME).
		SetHost(utils.OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/subchain/update")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("sp_sdk_ver", utils.SDKVERSION)
	client.AddGetParam("sp_sdk_lang", utils.SDKLANG)
	client.AddPostParam("subchain_id", params.SubchainID)
	client.AddPostParam("chain_name", params.ChainName)
	client.AddPostParam("chain_desc", params.ChainDesc)
	client.AddPostParam("chain_path", params.ChainPath)
	client.AddPostParam("telephone", params.Telephone)

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

package search

import (
	"github.com/baidu/smartapp-openapi-go/utils"
)

// RankSubChainRequest 请求结构体
type RankSubChainRequest struct {
	AccessToken      string // 授权小程序的接口调用凭据
	SubchainRanklist string // 子链 Id 字符串，顺序代表了排序位置,使用逗号分割
}

// 响应结构体

type RankSubChainResponse struct {
	Data      bool   `json:"data"`       // 是否成功
	Errno     int64  `json:"errno"`      // 状态码
	ErrMsg    string `json:"msg"`        // 错误信息
	ErrorCode int64  `json:"error_code"` // openapi 错误码
	ErrorMsg  string `json:"error_msg"`  // openapi 错误信息
}

// RankSubChain
func RankSubChain(params *RankSubChainRequest) (bool, error) {
	var (
		err        error
		defaultRet bool
	)
	respData := &RankSubChainResponse{}

	client := utils.NewHTTPClient().
		SetContentType(utils.ContentTypeForm).
		SetConverterType(utils.ConverterTypeJSON).
		SetMethod("POST").
		SetScheme(utils.SCHEME).
		SetHost(utils.OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/subchain/rank")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("sp_sdk_ver", utils.SDKVERSION)
	client.AddGetParam("sp_sdk_lang", utils.SDKLANG)
	client.AddPostParam("subchain_ranklist", params.SubchainRanklist)

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

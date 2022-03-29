package search

import (
	"fmt"

	"github.com/baidu/smartapp-openapi-go/utils"
)

// GetBindSiteRequest 请求结构体
type GetBindSiteRequest struct {
	AccessToken string // 授权小程序的接口调用凭据
	PageNum     int64  // 分页页码。从1开始
	PageSize    int64  // 分页大小。最大值为 100
}

// 响应结构体

type GetBindSiteResponse struct {
	Data      []string `json:"data"`       // 响应参数
	Errno     int64    `json:"errno"`      // 状态码
	ErrMsg    string   `json:"msg"`        // 错误信息
	ErrorCode int64    `json:"error_code"` // openapi 错误码
	ErrorMsg  string   `json:"error_msg"`  // openapi 错误信息
}

// GetBindSite
func GetBindSite(params *GetBindSiteRequest) ([]string, error) {
	var (
		err        error
		defaultRet []string
	)
	respData := &GetBindSiteResponse{}

	client := utils.NewHTTPClient().
		SetContentType(utils.ContentTypeForm).
		SetConverterType(utils.ConverterTypeJSON).
		SetMethod("POST").
		SetScheme(utils.SCHEME).
		SetHost(utils.OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/flow/getbindsite")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("sp_sdk_ver", utils.SDKVERSION)
	client.AddGetParam("sp_sdk_lang", utils.SDKLANG)
	client.AddPostParam("page_num", fmt.Sprintf("%v", params.PageNum))
	client.AddPostParam("page_size", fmt.Sprintf("%v", params.PageSize))

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

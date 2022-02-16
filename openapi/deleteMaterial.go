package openapi

import (
	"fmt"
)

// DeleteMaterialRequest 请求结构体
type DeleteMaterialRequest struct {
	AccessToken string // 接口调用凭证
	AppID       int64  // app_id
	ID          int64  // 物料 id ，添加物料时返回 id
	Path        string // 智能小程序内页链接，取值为添加物料时返回的 path
}

// 响应结构体

type DeleteMaterialResponse struct {
	Data      bool   `json:"data"`       // true：代表修改成功，false：代码修改失败
	Errno     int64  `json:"errno"`      // 错误码
	ErrMsg    string `json:"msg"`        // 错误信息
	ErrorCode int64  `json:"error_code"` // openapi 错误码
	ErrorMsg  string `json:"error_msg"`  // openapi 错误信息
}

// DeleteMaterial
func DeleteMaterial(params *DeleteMaterialRequest) (bool, error) {
	var (
		err        error
		defaultRet bool
	)
	respData := &DeleteMaterialResponse{}

	client := NewHTTPClient().
		SetContentType(ContentTypeForm).
		SetConverterType(ConverterTypeJSON).
		SetMethod("POST").
		SetScheme(SCHEME).
		SetHost(OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/articlemount/material/delete")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("sp_sdk_ver", SDKVERSION)
	client.AddGetParam("sp_sdk_lang", SDKLANG)
	client.AddPostParam("app_id", fmt.Sprintf("%v", params.AppID))
	client.AddPostParam("id", fmt.Sprintf("%v", params.ID))
	client.AddPostParam("path", params.Path)

	err = client.Do()
	if err != nil {
		return defaultRet, err
	}
	err = client.Convert(respData)
	if err != nil {
		return defaultRet, err
	}
	if respData.ErrorCode != 0 {
		return defaultRet, &OpenAPIError{respData.ErrorCode, respData.ErrorMsg, respData}
	}
	if respData.Errno != 0 {
		return defaultRet, &APIError{respData.Errno, respData.ErrMsg, respData}
	}

	return respData.Data, nil
}

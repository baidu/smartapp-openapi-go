package openapi

import (
	"fmt"
)

// DeleteSkuCouponRequest 请求结构体
type DeleteSkuCouponRequest struct {
	AccessToken string // 接口调用凭证
	AppID       int64  // app_id
	PathList    string // 需要删除的资源 path 列表
}

// 响应结构体

type DeleteSkuCouponResponse struct {
	Data      string `json:"data"`       // 响应对象
	Errno     int64  `json:"errno"`      // 错误码
	ErrMsg    string `json:"msg"`        // 错误信息
	ErrorCode int64  `json:"error_code"` // openapi 错误码
	ErrorMsg  string `json:"error_msg"`  // openapi 错误信息
}

// DeleteSkuCoupon
func DeleteSkuCoupon(params *DeleteSkuCouponRequest) (string, error) {
	var (
		err        error
		defaultRet string
	)
	respData := &DeleteSkuCouponResponse{}

	client := NewHTTPClient().
		SetContentType(ContentTypeForm).
		SetConverterType(ConverterTypeJSON).
		SetMethod("POST").
		SetScheme(SCHEME).
		SetHost(OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/server/delete/skuCoupon")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("app_id", fmt.Sprintf("%v", params.AppID))
	client.AddGetParam("path_list", params.PathList)
	client.AddGetParam("sp_sdk_ver", SDKVERSION)
	client.AddGetParam("sp_sdk_lang", SDKLANG)

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

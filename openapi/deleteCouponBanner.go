package openapi

import (
	"fmt"
)

// DeleteCouponBannerRequest 请求结构体
type DeleteCouponBannerRequest struct {
	AccessToken string // 接口调用凭证
	CouponID    string // 卡券 ID
	BannerID    int64  // 卡券 banner 记录 id
}

// 响应结构体

type DeleteCouponBannerResponse struct {
	Data      bool   `json:"data"`       // true 退还成功 false 退还失败
	Errno     int64  `json:"errno"`      // 错误码
	ErrMsg    string `json:"msg"`        // 错误信息
	ErrorCode int64  `json:"error_code"` // openapi 错误码
	ErrorMsg  string `json:"error_msg"`  // openapi 错误信息
}

// DeleteCouponBanner
func DeleteCouponBanner(params *DeleteCouponBannerRequest) (bool, error) {
	var (
		err        error
		defaultRet bool
	)
	respData := &DeleteCouponBannerResponse{}

	client := NewHTTPClient().
		SetContentType(ContentTypeForm).
		SetConverterType(ConverterTypeJSON).
		SetMethod("POST").
		SetScheme(SCHEME).
		SetHost(OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/v1.0/coupon/banner/delete")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("sp_sdk_ver", SDKVERSION)
	client.AddGetParam("sp_sdk_lang", SDKLANG)
	client.AddPostParam("couponId", params.CouponID)
	client.AddPostParam("bannerId", fmt.Sprintf("%v", params.BannerID))

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

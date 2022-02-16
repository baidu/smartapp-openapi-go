package openapi

import (
	"encoding/json"
)

// AddCouponBannerRequest 请求结构体
type AddCouponBannerRequest struct {
	AccessToken     string // 接口调用凭证
	CouponID        string // 卡券 ID
	PicURL          string // 卡券 banner 图片
	Title           string // 卡券 banner 图标题
	AppRedirectPath string // banner 图跳转的小程序页面路径
}

// 响应结构体

type AddCouponBannerResponsedata struct {
	BannerID int64 `json:"bannerId"` // 卡券 banner 记录 id
}

type AddCouponBannerResponse struct {
	Data      AddCouponBannerResponsedata `json:"data"`       // 响应对象
	Errno     int64                       `json:"errno"`      // 错误码
	ErrMsg    string                      `json:"msg"`        // 错误信息
	ErrorCode int64                       `json:"error_code"` // openapi 错误码
	ErrorMsg  string                      `json:"error_msg"`  // openapi 错误信息
}

// AddCouponBanner
func AddCouponBanner(params *AddCouponBannerRequest) (*AddCouponBannerResponsedata, error) {
	var (
		err        error
		defaultRet *AddCouponBannerResponsedata
	)
	respData := &AddCouponBannerResponse{}

	client := NewHTTPClient().
		SetContentType(ContentTypeJSON).
		SetConverterType(ConverterTypeJSON).
		SetMethod("POST").
		SetScheme(SCHEME).
		SetHost(OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/v1.0/coupon/banner/add")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("sp_sdk_ver", SDKVERSION)
	client.AddGetParam("sp_sdk_lang", SDKLANG)
	postData := map[string]interface{}{
		"couponId":        params.CouponID,
		"picUrl":          params.PicURL,
		"title":           params.Title,
		"appRedirectPath": params.AppRedirectPath,
	}
	bts, err := json.Marshal(postData)
	if err != nil {
		return defaultRet, err
	}
	client.SetBody(bts)

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

	return &respData.Data, nil
}

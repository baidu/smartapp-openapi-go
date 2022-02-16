package openapi


// SendCouponRequest 请求结构体
type SendCouponRequest struct {
	AccessToken           string // 接口调用凭证
	CouponID              string // 卡券 ID
	OpenID                string // 用户 ID
	UniqueSendingIdentity string // 请求唯一标识 uuid
}

// 响应结构体

type SendCouponResponsedata struct {
	CouponTakeID int64 `json:"couponTakeId"` // 卡券领取标识
}

type SendCouponResponse struct {
	Data      SendCouponResponsedata `json:"data"`       // 响应对象
	Errno     int64                  `json:"errno"`      // 错误码
	ErrMsg    string                 `json:"msg"`        // 错误信息
	ErrorCode int64                  `json:"error_code"` // openapi 错误码
	ErrorMsg  string                 `json:"error_msg"`  // openapi 错误信息
}

// SendCoupon
func SendCoupon(params *SendCouponRequest) (*SendCouponResponsedata, error) {
	var (
		err        error
		defaultRet *SendCouponResponsedata
	)
	respData := &SendCouponResponse{}

	client := NewHTTPClient().
		SetContentType(ContentTypeForm).
		SetConverterType(ConverterTypeJSON).
		SetMethod("POST").
		SetScheme(SCHEME).
		SetHost(OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/v1.0/coupon/send")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("sp_sdk_ver", SDKVERSION)
	client.AddGetParam("sp_sdk_lang", SDKLANG)
	client.AddPostParam("couponId", params.CouponID)
	client.AddPostParam("openId", params.OpenID)
	client.AddPostParam("uniqueSendingIdentity", params.UniqueSendingIdentity)

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

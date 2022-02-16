package openapi

// BatchUploadCouponCodeRequest 请求结构体
type BatchUploadCouponCodeRequest struct {
	AccessToken string // 接口调用凭证
	CouponID    string // 卡券 ID
	CouponCodes string // 卡券 Code 码列表
}

// 响应结构体

type BatchUploadCouponCodeResponsedata struct {
	FailNum    string `json:"failNum"`    // 请求失败数量
	SuccessNum string `json:"successNum"` // 请求成功数量
}

type BatchUploadCouponCodeResponse struct {
	Data      BatchUploadCouponCodeResponsedata `json:"data"`       // 响应对象
	Errno     int64                             `json:"errno"`      // 错误码
	ErrMsg    string                            `json:"msg"`        // 错误信息
	ErrorCode int64                             `json:"error_code"` // openapi 错误码
	ErrorMsg  string                            `json:"error_msg"`  // openapi 错误信息
}

// BatchUploadCouponCode
func BatchUploadCouponCode(params *BatchUploadCouponCodeRequest) (*BatchUploadCouponCodeResponsedata, error) {
	var (
		err        error
		defaultRet *BatchUploadCouponCodeResponsedata
	)
	respData := &BatchUploadCouponCodeResponse{}

	client := NewHTTPClient().
		SetContentType(ContentTypeForm).
		SetConverterType(ConverterTypeJSON).
		SetMethod("POST").
		SetScheme(SCHEME).
		SetHost(OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/v1.0/coupon/code/batchUpload")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("sp_sdk_ver", SDKVERSION)
	client.AddGetParam("sp_sdk_lang", SDKLANG)
	client.AddPostParam("couponId", params.CouponID)
	client.AddPostParam("couponCodes", params.CouponCodes)

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

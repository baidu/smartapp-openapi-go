package openapi


// GetCouponBannerRequest 请求结构体
type GetCouponBannerRequest struct {
	AccessToken string // 接口调用凭证
	CouponID    string // 卡券 ID
	BannerIds   string // 卡券 banner 记录 id
}

// 响应结构体

type GetCouponBannerResponsedataItem struct {
	AppRedirectPath string `json:"appRedirectPath"` // banner 图跳转的小程序页面路径
	BannerID        int64  `json:"bannerId"`        // 卡券 banner 记录 id
	CouponID        string `json:"couponId"`        // 卡券 ID
	CreateTime      int64  `json:"createTime"`      // 卡券创建时间
	PicURL          string `json:"picUrl"`          // 卡券 banner 图片
	Title           string `json:"title"`           // 卡券 banner 图标题
	UpdateTime      int64  `json:"updateTime"`      // 卡券更新事件
}

type GetCouponBannerResponse struct {
	Data      []GetCouponBannerResponsedataItem `json:"data"`       // 卡券数组
	Errno     int64                             `json:"errno"`      // 错误码
	ErrMsg    string                            `json:"msg"`        // 错误信息
	ErrorCode int64                             `json:"error_code"` // openapi 错误码
	ErrorMsg  string                            `json:"error_msg"`  // openapi 错误信息
}

// GetCouponBanner
func GetCouponBanner(params *GetCouponBannerRequest) ([]GetCouponBannerResponsedataItem, error) {
	var (
		err        error
		defaultRet []GetCouponBannerResponsedataItem
	)
	respData := &GetCouponBannerResponse{}

	client := NewHTTPClient().
		SetContentType(ContentTypeForm).
		SetConverterType(ConverterTypeJSON).
		SetMethod("GET").
		SetScheme(SCHEME).
		SetHost(OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/v1.0/coupon/banner/get")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("couponId", params.CouponID)
	client.AddGetParam("bannerIds", params.BannerIds)
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

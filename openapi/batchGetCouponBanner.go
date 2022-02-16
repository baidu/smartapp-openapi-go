package openapi

// BatchGetCouponBannerRequest 请求结构体
type BatchGetCouponBannerRequest struct {
	AccessToken string // 接口调用凭证
	CouponID    string // 卡券 ID
}

// 响应结构体

type BatchGetCouponBannerResponsedatadataListItem struct {
	AppRedirectPath string `json:"appRedirectPath"` // banner 图跳转的小程序页面路径
	BannerID        int64  `json:"bannerId"`        // 卡券 banner 记录 id
	CouponID        string `json:"couponId"`        // 卡券 id
	CreateTime      int64  `json:"createTime"`      // 卡券创建时间
	PicURL          string `json:"picUrl"`          // 卡券图片推广图链接地址
	Title           string `json:"title"`           // 卡券图片推广图标题
	UpdateTime      int64  `json:"updateTime"`      // 卡券更新时间
}

type BatchGetCouponBannerResponsedata struct {
	DataList []BatchGetCouponBannerResponsedatadataListItem `json:"dataList"` // 响应卡券数据数组
	PageNo   int64                                          `json:"pageNo"`   // 响应数据页数
	Total    int64                                          `json:"total"`    // 响应数据量
}

type BatchGetCouponBannerResponse struct {
	Data      BatchGetCouponBannerResponsedata `json:"data"`       // 响应对象
	Errno     int64                            `json:"errno"`      // 错误码
	ErrMsg    string                           `json:"msg"`        // 错误信息
	ErrorCode int64                            `json:"error_code"` // openapi 错误码
	ErrorMsg  string                           `json:"error_msg"`  // openapi 错误信息
}

// BatchGetCouponBanner
func BatchGetCouponBanner(params *BatchGetCouponBannerRequest) (*BatchGetCouponBannerResponsedata, error) {
	var (
		err        error
		defaultRet *BatchGetCouponBannerResponsedata
	)
	respData := &BatchGetCouponBannerResponse{}

	client := NewHTTPClient().
		SetContentType(ContentTypeForm).
		SetConverterType(ConverterTypeJSON).
		SetMethod("GET").
		SetScheme(SCHEME).
		SetHost(OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/v1.0/coupon/banner/batchGet")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("couponId", params.CouponID)
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

	return &respData.Data, nil
}

package openapi

// GetUnionidRequest 请求结构体
type GetUnionidRequest struct {
	AccessToken string // 接口调用凭证
	Openid      string // 用户 openid ，需要经过用户登录授权过程获取
}

// 响应结构体

type GetUnionidResponsedata struct {
	Unionid string `json:"unionid"` //
}

type GetUnionidResponse struct {
	Data      GetUnionidResponsedata `json:"data"`       // 详细数据，errno 为 0 的情况下才有意义
	ErrMsg    string                 `json:"msg"`        // 错误信息
	Errno     int64                  `json:"errno"`      // 错误码
	RequestID string                 `json:"request_id"` // 请求 ID ，标识一次请求
	Timestamp int64                  `json:"timestamp"`  // 请求时间戳
	ErrorCode int64                  `json:"error_code"` // openapi 错误码
	ErrorMsg  string                 `json:"error_msg"`  // openapi 错误信息
}

// GetUnionid
func GetUnionid(params *GetUnionidRequest) (*GetUnionidResponsedata, error) {
	var (
		err        error
		defaultRet *GetUnionidResponsedata
	)
	respData := &GetUnionidResponse{}

	client := NewHTTPClient().
		SetContentType(ContentTypeForm).
		SetConverterType(ConverterTypeJSON).
		SetMethod("POST").
		SetScheme(SCHEME).
		SetHost(OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/getunionid")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("sp_sdk_ver", SDKVERSION)
	client.AddGetParam("sp_sdk_lang", SDKLANG)
	client.AddPostParam("openid", params.Openid)

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

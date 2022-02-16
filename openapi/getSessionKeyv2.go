package openapi


// GetSessionKeyv2Request 请求结构体
type GetSessionKeyv2Request struct {
	AccessToken string // 接口调用凭证
	Code        string // 通过 swan.getLoginCode 获取 Authorization Code 特殊说明：code 中有 @ 符号时，会请求对应的开源宿主，用户身份校验及 SessionKey 生成过程由开源宿主实现
}

// 响应结构体

type GetSessionKeyv2Responsedata struct {
	OpenID     string `json:"open_id"`     // 用户身份标识 不同用户登录同一个小程序获取到的 openid 不同，同一个用户登录不同小程序获取到的 openid 也不同
	SessionKey string `json:"session_key"` // 用户的 SessionKey
}

type GetSessionKeyv2Response struct {
	Data      GetSessionKeyv2Responsedata `json:"data"`       // 响应对象
	ErrMsg    string                      `json:"msg"`        // 错误信息
	Errno     int64                       `json:"errno"`      // 错误码
	RequestID string                      `json:"request_id"` // 请求 id，标识一次请求
	Timestamp int64                       `json:"timestamp"`  // 时间戳
	ErrorCode int64                       `json:"error_code"` // openapi 错误码
	ErrorMsg  string                      `json:"error_msg"`  // openapi 错误信息
}

// GetSessionKeyv2
func GetSessionKeyv2(params *GetSessionKeyv2Request) (*GetSessionKeyv2Responsedata, error) {
	var (
		err        error
		defaultRet *GetSessionKeyv2Responsedata
	)
	respData := &GetSessionKeyv2Response{}

	client := NewHTTPClient().
		SetContentType(ContentTypeForm).
		SetConverterType(ConverterTypeJSON).
		SetMethod("POST").
		SetScheme(SCHEME).
		SetHost(OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/getsessionkey")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("code", params.Code)
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

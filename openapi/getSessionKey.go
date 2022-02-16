package openapi


// GetSessionKeyRequest 请求结构体
type GetSessionKeyRequest struct {
	Code     string // 通过 swan.getLoginCode 获取 Authorization Code 特殊说明：code 中有 @ 符号时，会请求对应的开源宿主，用户身份校验及 SessionKey 生成过程由开源宿主实现
	ClientID string // 智能小程序的 AppKey
	Sk       string // 智能小程序的 AppSecret
}

// 响应结构体
type GetSessionKeyResponsedata struct {
	Openid     string
	SessionKey string
}

type GetSessionKeyResponse struct {
	Errno            int64  `json:"errno"`             // 错误码，详情见下方错误码
	ErrMsg           string `json:"msg"`               // 错误信息
	ErrorDescription string `json:"error_description"` // 错误描述信息，用来帮助理解和解决发生的错误
	Openid           string `json:"openid"`            // 用户身份标识 不同用户登录同一个小程序获取到的 openid 不同，同一个用户登录不同小程序获取到的 openid 也不同
	SessionKey       string `json:"session_key"`       // 用户的 SessionKey
	ErrorCode        int64  `json:"error_code"`        // openapi 错误码
	ErrorMsg         string `json:"error_msg"`         // openapi 错误信息
}

// GetSessionKey
func GetSessionKey(params *GetSessionKeyRequest) (*GetSessionKeyResponsedata, error) {
	var (
		err        error
		defaultRet *GetSessionKeyResponsedata
	)
	respData := &GetSessionKeyResponse{}

	client := NewHTTPClient().
		SetContentType(ContentTypeForm).
		SetConverterType(ConverterTypeJSON).
		SetMethod("GET").
		SetScheme(SCHEME).
		SetHost(SPAPIHOST).
		SetPath("/oauth/jscode2sessionkey")
	client.AddGetParam("code", params.Code)
	client.AddGetParam("client_id", params.ClientID)
	client.AddGetParam("sk", params.Sk)
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

	resData := &GetSessionKeyResponsedata{
		Openid:     respData.Openid,
		SessionKey: respData.SessionKey,
	}

	return resData, nil
}

package openapi


// GetAccessTokenRequest 请求结构体
type GetAccessTokenRequest struct {
	GrantType    string // 固定为：client_credentials
	ClientID     string // 智能小程序的 App Key，可在「开发者平台 - 设置 - 开发设置」页中获得。（要先在开发者平台创建小程序）
	ClientSecret string // 智能小程序的 App Secret，请妥善保存，如丢失可在「开发者平台 - 设置 - 开发设置」页面重置后获得，重置后 App Secret 将会被更新。
	Scope        string // 固定为：smartapp_snsapi_base
}

// 响应结构体
type GetAccessTokenResponsedata struct {
	AccessToken string
	ExpiresIn   int64
}

type GetAccessTokenResponse struct {
	AccessToken      string `json:"access_token"`      // 获取到的凭证
	Error            string `json:"error"`             // 异常提示信息
	ErrorDescription string `json:"error_description"` // 异常情况详细的提示信息
	ExpiresIn        int64  `json:"expires_in"`        // 凭证有效时间，单位：秒
	ErrorCode        int64  `json:"error_code"`        // openapi 错误码
	ErrorMsg         string `json:"error_msg"`         // openapi 错误信息
}

// GetAccessToken
func GetAccessToken(params *GetAccessTokenRequest) (*GetAccessTokenResponsedata, error) {
	var (
		err        error
		defaultRet *GetAccessTokenResponsedata
	)
	respData := &GetAccessTokenResponse{}

	client := NewHTTPClient().
		SetContentType(ContentTypeForm).
		SetConverterType(ConverterTypeJSON).
		SetMethod("GET").
		SetScheme(SCHEME).
		SetHost(OPENAPIHOST).
		SetPath("/oauth/2.0/token")
	client.AddGetParam("grant_type", params.GrantType)
	client.AddGetParam("client_id", params.ClientID)
	client.AddGetParam("client_secret", params.ClientSecret)
	client.AddGetParam("scope", params.Scope)
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
	if respData.Error != "" {
		return defaultRet, &OpenAPIError{-1, respData.Error, respData}
	}

	resData := &GetAccessTokenResponsedata{
		AccessToken: respData.AccessToken,
		ExpiresIn:   respData.ExpiresIn,
	}

	return resData, nil
}

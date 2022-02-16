package openapi


// GetLikeCountRequest 请求结构体
type GetLikeCountRequest struct {
	AccessToken string // 接口调用凭证
	HostName    string // 宿主名称
	Snid        string // 文章的 ID
	SnidType    string // 无 snid_type 的开发者请默认传空串
}

// 响应结构体

type GetLikeCountResponsedata struct {
	LikeCount int64 `json:"like_count"` // 点赞数
}

type GetLikeCountResponse struct {
	Data      GetLikeCountResponsedata `json:"data"`       // 响应对象
	Errno     int64                    `json:"errno"`      // 错误码
	ErrMsg    string                   `json:"msg"`        // 错误信息
	RequestID string                   `json:"request_id"` // 请求 ID，标识一次请求
	ErrorCode int64                    `json:"error_code"` // openapi 错误码
	ErrorMsg  string                   `json:"error_msg"`  // openapi 错误信息
}

// GetLikeCount
func GetLikeCount(params *GetLikeCountRequest) (*GetLikeCountResponsedata, error) {
	var (
		err        error
		defaultRet *GetLikeCountResponsedata
	)
	respData := &GetLikeCountResponse{}

	client := NewHTTPClient().
		SetContentType(ContentTypeForm).
		SetConverterType(ConverterTypeJSON).
		SetMethod("POST").
		SetScheme(SCHEME).
		SetHost(OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/oss/publisher/ugc/like_count")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("host_name", params.HostName)
	client.AddGetParam("sp_sdk_ver", SDKVERSION)
	client.AddGetParam("sp_sdk_lang", SDKLANG)
	client.AddPostParam("snid", params.Snid)
	client.AddPostParam("snid_type", params.SnidType)

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

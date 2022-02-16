package openapi


// GetCommentCountRequest 请求结构体
type GetCommentCountRequest struct {
	AccessToken string // 接口调用凭证
	Snid        string // 文章的 ID
	SnidType    string // 无 snid_type 的开发者请默认传空串
}

// 响应结构体

type GetCommentCountResponsedata struct {
	CommentCount string `json:"comment_count"` // 评论数
}

type GetCommentCountResponse struct {
	Data      GetCommentCountResponsedata `json:"data"`       // 互动评论数详细结果
	Errno     int64                       `json:"errno"`      // 错误码，正常返回情况下，errno 值为 0
	ErrMsg    string                      `json:"msg"`        // 错误信息
	RequestID string                      `json:"request_id"` // 请求 ID，标识一次请求
	ErrorCode int64                       `json:"error_code"` // openapi 错误码
	ErrorMsg  string                      `json:"error_msg"`  // openapi 错误信息
}

// GetCommentCount
func GetCommentCount(params *GetCommentCountRequest) (*GetCommentCountResponsedata, error) {
	var (
		err        error
		defaultRet *GetCommentCountResponsedata
	)
	respData := &GetCommentCountResponse{}

	client := NewHTTPClient().
		SetContentType(ContentTypeForm).
		SetConverterType(ConverterTypeJSON).
		SetMethod("POST").
		SetScheme(SCHEME).
		SetHost(OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/ma/component/comment/tcount")
	client.AddGetParam("access_token", params.AccessToken)
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

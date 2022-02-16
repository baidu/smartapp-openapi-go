package openapi


// GetPathCheckResultByIDRequest 请求结构体
type GetPathCheckResultByIDRequest struct {
	AccessToken string // 接口调用凭据
	RetrieveID  string // 异步检测页面接口返回的唯一标识
}

// 响应结构体

type GetPathCheckResultByIDResponsedatadetailItem struct {
	Errno   int64    `json:"errno"`   // 错误码
	HarmURL []string `json:"harmUrl"` // 违规图片的链接列表，图片相关的检测策略该字段可能会有值
	Msg     string   `json:"msg"`     // 错误信息
	Type    string   `json:"type"`    // 检测策略
}

type GetPathCheckResultByIDResponsedata struct {
	CheckStatus int64                                          `json:"checkStatus"` // 检测状态，0 为初始化，1 为内容抓取，2 为检测中，3 为检测完成，4 为检测失败
	CreateTime  int64                                          `json:"createTime"`  // 创建时间戳，单位毫秒
	Detail      []GetPathCheckResultByIDResponsedatadetailItem `json:"detail"`      // 是一个对象数组，返回每一种检测策略的结果
}

type GetPathCheckResultByIDResponse struct {
	Data      GetPathCheckResultByIDResponsedata `json:"data"`       // 响应对象
	Errno     int64                              `json:"errno"`      // 错误码
	ErrMsg    string                             `json:"msg"`        // 错误信息
	ErrorCode int64                              `json:"error_code"` // openapi 错误码
	ErrorMsg  string                             `json:"error_msg"`  // openapi 错误信息
}

// GetPathCheckResultByID
func GetPathCheckResultByID(params *GetPathCheckResultByIDRequest) (*GetPathCheckResultByIDResponsedata, error) {
	var (
		err        error
		defaultRet *GetPathCheckResultByIDResponsedata
	)
	respData := &GetPathCheckResultByIDResponse{}

	client := NewHTTPClient().
		SetContentType(ContentTypeForm).
		SetConverterType(ConverterTypeJSON).
		SetMethod("GET").
		SetScheme(SCHEME).
		SetHost(OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/riskDetection/v2/getPathCheckResultByID")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("retrieveId", params.RetrieveID)
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

package openapi

// MisjudgeFeedbackRequest 请求结构体
type MisjudgeFeedbackRequest struct {
	AccessToken string // 接口调用凭据
	RetrieveID  string // 前面提交检测接口返回的 retrieveId
}

// 响应结构体
type MisjudgeFeedbackResponsedata struct {
	Errno int64
	Msg   string
}

type MisjudgeFeedbackResponse struct {
	Errno     int64  `json:"errno"`      // 错误码
	ErrMsg    string `json:"msg"`        // 错误信息
	ErrorCode int64  `json:"error_code"` // openapi 错误码
	ErrorMsg  string `json:"error_msg"`  // openapi 错误信息
}

// MisjudgeFeedback
func MisjudgeFeedback(params *MisjudgeFeedbackRequest) (*MisjudgeFeedbackResponsedata, error) {
	var (
		err        error
		defaultRet *MisjudgeFeedbackResponsedata
	)
	respData := &MisjudgeFeedbackResponse{}

	client := NewHTTPClient().
		SetContentType(ContentTypeForm).
		SetConverterType(ConverterTypeJSON).
		SetMethod("POST").
		SetScheme(SCHEME).
		SetHost(OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/riskDetection/misjudgeFeedback")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("sp_sdk_ver", SDKVERSION)
	client.AddGetParam("sp_sdk_lang", SDKLANG)
	client.AddPostParam("retrieveId", params.RetrieveID)

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

	resData := &MisjudgeFeedbackResponsedata{
		Errno: respData.Errno,
		Msg:   respData.ErrMsg,
	}

	return resData, nil
}

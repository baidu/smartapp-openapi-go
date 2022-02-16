package openapi


// DeleteTemplateRequest 请求结构体
type DeleteTemplateRequest struct {
	AccessToken string // 接口调用凭证
	TemplateID  string // 模板 id ，发送小程序模板消息时所需
}

// 响应结构体
type DeleteTemplateResponsedata struct {
	Errno int64
	Msg   string
}

type DeleteTemplateResponse struct {
	Errno     int64  `json:"errno"`      // 错误码
	ErrMsg    string `json:"msg"`        // 错误信息
	ErrorCode int64  `json:"error_code"` // openapi 错误码
	ErrorMsg  string `json:"error_msg"`  // openapi 错误信息
}

// DeleteTemplate
func DeleteTemplate(params *DeleteTemplateRequest) (*DeleteTemplateResponsedata, error) {
	var (
		err        error
		defaultRet *DeleteTemplateResponsedata
	)
	respData := &DeleteTemplateResponse{}

	client := NewHTTPClient().
		SetContentType(ContentTypeForm).
		SetConverterType(ConverterTypeJSON).
		SetMethod("POST").
		SetScheme(SCHEME).
		SetHost(OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/template/templatedel")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("sp_sdk_ver", SDKVERSION)
	client.AddGetParam("sp_sdk_lang", SDKLANG)
	client.AddPostParam("template_id", params.TemplateID)

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

	resData := &DeleteTemplateResponsedata{
		Errno: respData.Errno,
		Msg:   respData.ErrMsg,
	}

	return resData, nil
}

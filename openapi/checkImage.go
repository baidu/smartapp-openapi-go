package openapi

import (
	"bytes"
	"io"
	"mime/multipart"
	"path/filepath"
	
)

// CheckImageRequest 请求结构体
type CheckImageRequest struct {
	AccessToken string                // 接口调用凭据
	Image       CheckImageRequestFile // 图片文件，只支持 PNG、JPG、JPEG 三种格式，且文件大小不能超过 5MB
	Type        string                // 检测策略，porn 为色情检测，ocr-word 为图片上文字的词表检测，ocr-lead 为图片上文字的诱导检测。可以多选，多个值之间用英文逗号拼接，不传默认为 porn，参数值区分大小写
}

type CheckImageRequestFile struct {
	Name   string
	Reader io.Reader
}

// 响应结构体

type CheckImageResponsedataresItem struct {
	Errno int64  `json:"errno"` // 错误码
	Msg   string `json:"msg"`   // 错误信息
	Type  string `json:"type"`  // 检测策略
}

type CheckImageResponsedata struct {
	Res        []CheckImageResponsedataresItem `json:"res"`        // 是一个对象数组，返回每一种检测策略的结果
	RetrieveID string                          `json:"retrieveId"` // 调用误判反馈接口时需要该返回值
}

type CheckImageResponse struct {
	Data      CheckImageResponsedata `json:"data"`       // 响应对象
	Errno     int64                  `json:"errno"`      // 错误码
	ErrMsg    string                 `json:"msg"`        // 错误信息
	ErrorCode int64                  `json:"error_code"` // openapi 错误码
	ErrorMsg  string                 `json:"error_msg"`  // openapi 错误信息
}

// CheckImage
func CheckImage(params *CheckImageRequest) (*CheckImageResponsedata, error) {
	var (
		err        error
		defaultRet *CheckImageResponsedata
	)
	respData := &CheckImageResponse{}

	// Post MultiPart/form-data
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)

	part1, err := writer.CreateFormFile(params.Image.Name, filepath.Base(params.Image.Name))
	if err != nil {
		return defaultRet, err
	}

	_, err = io.Copy(part1, params.Image.Reader)
	if err != nil {
		return defaultRet, err
	}
	_ = writer.WriteField("type", params.Type)

	err = writer.Close()
	if err != nil {
		return defaultRet, err
	}

	client := NewHTTPClient().
		SetContentType(ContentTypeMultiPart + "; boundary=" + writer.Boundary()).
		SetBody(payload).
		SetConverterType(ConverterTypeJSON).
		SetMethod("POST").
		SetScheme(SCHEME).
		SetHost(OPENAPIHOST).
		SetPath("/file/2.0/smartapp/riskDetection/v2/syncCheckImage")
	client.AddGetParam("access_token", params.AccessToken)
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

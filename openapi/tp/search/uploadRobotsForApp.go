package search

import (
	"bytes"
	"io"
	"mime/multipart"
	"path/filepath"

	"github.com/baidu/smartapp-openapi-go/utils"
)

// UploadRobotsForAppRequest 请求结构体
type UploadRobotsForAppRequest struct {
	AccessToken string                        // 授权小程序的接口调用凭据
	Robots      UploadRobotsForAppRequestFile // robots.txt 文件，要求文件为 txt 格式，目前支持48k的文件内容检测，请保证robots.txt文件不要过大，目录最长不超过250个字符。
}

type UploadRobotsForAppRequestFile struct {
	Name   string
	Reader io.Reader
}

// 响应结构体

type UploadRobotsForAppResponse struct {
	Data      string `json:"data"`       // 响应参数
	Errno     int64  `json:"errno"`      // 状态码
	ErrMsg    string `json:"msg"`        // 错误信息
	ErrorCode int64  `json:"error_code"` // openapi 错误码
	ErrorMsg  string `json:"error_msg"`  // openapi 错误信息
}

// UploadRobotsForApp
func UploadRobotsForApp(params *UploadRobotsForAppRequest) (string, error) {
	var (
		err        error
		defaultRet string
	)
	respData := &UploadRobotsForAppResponse{}

	// Post MultiPart/form-data
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)

	part1, err := writer.CreateFormFile(params.Robots.Name, filepath.Base(params.Robots.Name))
	if err != nil {
		return defaultRet, err
	}

	_, err = io.Copy(part1, params.Robots.Reader)
	if err != nil {
		return defaultRet, err
	}

	err = writer.Close()
	if err != nil {
		return defaultRet, err
	}

	client := utils.NewHTTPClient().
		SetContentType(utils.ContentTypeMultiPart + "; boundary=" + writer.Boundary()).
		SetBody(payload).
		SetConverterType(utils.ConverterTypeJSON).
		SetMethod("POST").
		SetScheme(utils.SCHEME).
		SetHost(utils.OPENAPIHOST).
		SetPath("/file/2.0/smartapp/robots/app/upload")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("sp_sdk_ver", utils.SDKVERSION)
	client.AddGetParam("sp_sdk_lang", utils.SDKLANG)

	err = client.Do()
	if err != nil {
		return defaultRet, err
	}
	err = client.Convert(respData)
	if err != nil {
		return defaultRet, err
	}
	if respData.ErrorCode != 0 {
		return defaultRet, &utils.OpenAPIError{respData.ErrorCode, respData.ErrorMsg, respData}
	}

	if respData.Errno != 0 {
		return defaultRet, &utils.APIError{respData.Errno, respData.ErrMsg, respData}
	}
	return respData.Data, nil
}

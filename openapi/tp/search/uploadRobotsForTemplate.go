package search

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"path/filepath"

	"github.com/baidu/smartapp-openapi-go/utils"
)

// UploadRobotsForTemplateRequest 请求结构体
type UploadRobotsForTemplateRequest struct {
	AccessToken string                             // 授权小程序的接口调用凭据
	TemplateID  int64                              // 模版 id
	Robots      UploadRobotsForTemplateRequestFile // robots.txt 文件，要求文件为 txt 格式，目前支持 48k 的文件内容检测，请保证 robots.txt 文件不要过大，目录最长不超过 250 个字符
}

type UploadRobotsForTemplateRequestFile struct {
	Name   string
	Reader io.Reader
}

// 响应结构体

type UploadRobotsForTemplateResponse struct {
	Data      string `json:"data"`       // 响应参数
	Errno     int64  `json:"errno"`      // 状态码
	ErrMsg    string `json:"msg"`        // 错误信息
	ErrorCode int64  `json:"error_code"` // openapi 错误码
	ErrorMsg  string `json:"error_msg"`  // openapi 错误信息
}

// UploadRobotsForTemplate
func UploadRobotsForTemplate(params *UploadRobotsForTemplateRequest) (string, error) {
	var (
		err        error
		defaultRet string
	)
	respData := &UploadRobotsForTemplateResponse{}

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
		SetPath("/file/2.0/smartapp/robots/template/upload")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("template_id", fmt.Sprintf("%v", params.TemplateID))
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

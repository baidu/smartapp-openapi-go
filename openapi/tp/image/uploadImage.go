package image

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"path/filepath"

	"github.com/baidu/smartapp-openapi-go/utils"
)

// UploadImageRequest 请求结构体
type UploadImageRequest struct {
	AccessToken   string                 // 第三方平台的接口调用凭据
	MultipartFile UploadImageRequestFile // 文件
	Type          interface{}            // 图片用途，1/null：小程序头像；2：服务类目资质图片；3：小程序名称审核资料
}

type UploadImageRequestFile struct {
	Name   string
	Reader io.Reader
}

// 响应结构体

type UploadImageResponse struct {
	Data      string `json:"data"`       // 响应参数
	Errno     int64  `json:"errno"`      // 状态码
	ErrMsg    string `json:"msg"`        // 错误信息
	ErrorCode int64  `json:"error_code"` // openapi 错误码
	ErrorMsg  string `json:"error_msg"`  // openapi 错误信息
}

// UploadImage
func UploadImage(params *UploadImageRequest) (string, error) {
	var (
		err        error
		defaultRet string
	)
	respData := &UploadImageResponse{}

	// Post MultiPart/form-data
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)

	part1, err := writer.CreateFormFile(params.MultipartFile.Name, filepath.Base(params.MultipartFile.Name))
	if err != nil {
		return defaultRet, err
	}

	_, err = io.Copy(part1, params.MultipartFile.Reader)
	if err != nil {
		return defaultRet, err
	}
	_ = writer.WriteField("type", fmt.Sprint(params.Type))

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
		SetPath("/file/2.0/smartapp/upload/image")
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

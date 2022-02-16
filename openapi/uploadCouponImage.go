package openapi

import (
	"bytes"
	"io"
	"mime/multipart"
	"path/filepath"
)

// UploadCouponImageRequest 请求结构体
type UploadCouponImageRequest struct {
	AccessToken string                       // 接口调用凭证
	File        UploadCouponImageRequestFile // 卡券图
}

type UploadCouponImageRequestFile struct {
	Name   string
	Reader io.Reader
}

// 响应结构体

type UploadCouponImageResponsedata struct {
	URL string `json:"url"` // 图片地址
}

type UploadCouponImageResponse struct {
	Data      UploadCouponImageResponsedata `json:"data"`       // 响应对象
	Errno     int64                         `json:"errno"`      // 错误码
	ErrMsg    string                        `json:"msg"`        // 错误信息
	ErrorCode int64                         `json:"error_code"` // openapi 错误码
	ErrorMsg  string                        `json:"error_msg"`  // openapi 错误信息
}

// UploadCouponImage
func UploadCouponImage(params *UploadCouponImageRequest) (*UploadCouponImageResponsedata, error) {
	var (
		err        error
		defaultRet *UploadCouponImageResponsedata
	)
	respData := &UploadCouponImageResponse{}

	// Post MultiPart/form-data
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)

	part1, err := writer.CreateFormFile(params.File.Name, filepath.Base(params.File.Name))
	if err != nil {
		return defaultRet, err
	}

	_, err = io.Copy(part1, params.File.Reader)
	if err != nil {
		return defaultRet, err
	}

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
		SetPath("/file/2.0/smartapp/v1.0/coupon/upload/image")
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

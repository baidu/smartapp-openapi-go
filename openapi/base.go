package openapi

import "encoding/json"

const (
	SCHEME      = "https"
	OPENAPIHOST = "openapi.baidu.com"
	SPAPIHOST   = "spapi.baidu.com"
	SDKLANG     = "Go"
	SDKVERSION  = "1.0.8"
)

// 网关错误响应结构体
type OpenAPIError struct {
	ErrorCode int64       // openapi 错误码
	ErrorMsg  string      // openapi 错误信息
	Detail    interface{} // 详情
}

func (e *OpenAPIError) Error() string {
	bts, err := json.Marshal(e.Detail)
	if err != nil {
		return err.Error()
	}
	return string(bts)
}

// 业务接口错误响应结构体
type APIError struct {
	ErrNo  int64       // 错误码
	ErrMsg string      // 错误信息
	Detail interface{} // 详情
}

func (e *APIError) Error() string {
	bts, err := json.Marshal(e.Detail)
	if err != nil {
		return err.Error()
	}
	return string(bts)
}

package openapi

import (
	"fmt"
)

// DetectRiskRequest 请求结构体
type DetectRiskRequest struct {
	AccessToken string // 接口调用凭证
	Appkey      string // 小程序 appkey，智能小程序 AppKey
	Xtoken      string // 小程序通过swan-getSystemRiskInfo获取的内容,格式：{"key":"xxxx","value":"xxxx"}
	Type        string // 运营活动的类型，该值由风控平台分配。目前只有一种 marketing
	Clientip    string // 客户端的 IP，非小程序服务本地 IP
	Ts          int64  // 服务器的时间戳，秒级别
	Ev          string // 事件类型，预先分配事件 ID 定义。1、点击活动按钮（或者活动操作），活动相关操作默认选择此事件；2、 进入活动页面；3、注册；4、登录；5、分享；6、点赞；7、评论；8、 提现；9、下单/提单；10、支付；11、业务自定义动作；12、浏览 feed；13、开宝箱；14、领取红包；15、分享 feed；16、做任务；17、签到；18、排行榜；19、邀请；20、新客红包；21、摇一摇；22、语音红包；23、视频红包；24、金融授信；25、答题
	Useragent   string // 客户端请求小程序 Server 的 useragent，示例：Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.75 Safari/537.36
	Phone       string // 加密后的电话号码，加密方法：sha1
}

// 响应结构体

type DetectRiskResponsedata struct {
	Level string   `json:"level"` // 风险级别，目前有四个，风险等级依次降低（其中 1 最高，4 最低) 1 - 高危；2 - 嫌疑；3 - 普通；4 - 正常。 建议开发者将风险等级为 1、2 的做拦截处理
	Tag   []string `json:"tag"`   // 对应的描述
}

type DetectRiskResponse struct {
	Data      DetectRiskResponsedata `json:"data"`       // 响应对象
	ErrMsg    string                 `json:"msg"`        // 错误信息
	Errno     int64                  `json:"errno"`      // 错误码 0：正确返回，非 0：异常返回
	RequestID string                 `json:"request_id"` // 请求 ID，标识一次请求
	Timestamp int64                  `json:"timestamp"`  // 时间戳
	ErrorCode int64                  `json:"error_code"` // openapi 错误码
	ErrorMsg  string                 `json:"error_msg"`  // openapi 错误信息
}

// DetectRisk
func DetectRisk(params *DetectRiskRequest) (*DetectRiskResponsedata, error) {
	var (
		err        error
		defaultRet *DetectRiskResponsedata
	)
	respData := &DetectRiskResponse{}

	client := NewHTTPClient().
		SetContentType(ContentTypeForm).
		SetConverterType(ConverterTypeJSON).
		SetMethod("POST").
		SetScheme(SCHEME).
		SetHost(OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/detectrisk")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("sp_sdk_ver", SDKVERSION)
	client.AddGetParam("sp_sdk_lang", SDKLANG)
	client.AddPostParam("appkey", params.Appkey)
	client.AddPostParam("xtoken", params.Xtoken)
	client.AddPostParam("type", params.Type)
	client.AddPostParam("clientip", params.Clientip)
	client.AddPostParam("ts", fmt.Sprintf("%v", params.Ts))
	client.AddPostParam("ev", params.Ev)
	client.AddPostParam("useragent", params.Useragent)
	client.AddPostParam("phone", params.Phone)

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

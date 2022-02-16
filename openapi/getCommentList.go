package openapi

import (
	"fmt"
)

// GetCommentListRequest 请求结构体
type GetCommentListRequest struct {
	AccessToken string // 接口调用凭证
	HostName    string // 宿主名称
	Snid        string // 文章的 ID
	SnidType    string // 无 snid_type 的开发者请默认传空串
	Start       int64  // 评论的开始偏移量，默认0
	Num         int64  // 获取的评论条数，默认10
}

// 响应结构体

type GetCommentListResponsedatalistItemreplyListItem struct {
	Content    string `json:"content"`     // 回复内容
	CreateTime string `json:"create_time"` // 发表时间
	IsUped     string `json:"is_uped"`     // 点赞状态 1:已点赞 0:未点赞
	LikeCount  string `json:"like_count"`  // 点赞数
}

type GetCommentListResponsedatalistItem struct {
	Content    string                                            `json:"content"`     // 评论内容
	CreateTime string                                            `json:"create_time"` // 发表时间
	IsUped     string                                            `json:"is_uped"`     // 点赞状态 1:已点赞 0:未点赞
	LikeCount  string                                            `json:"like_count"`  // 点赞数
	ReplyList  []GetCommentListResponsedatalistItemreplyListItem `json:"reply_list"`  // 回复列表
	Srid       string                                            `json:"srid"`        // 评论ID
	ThreadID   string                                            `json:"thread_id"`   // 主题ID
}

type GetCommentListResponsedata struct {
	List []GetCommentListResponsedatalistItem `json:"list"` // 评论列表
}

type GetCommentListResponse struct {
	Data      GetCommentListResponsedata `json:"data"`       // 响应对象
	Errno     int64                      `json:"errno"`      // 错误码
	ErrMsg    string                     `json:"msg"`        // 错误信息
	RequestID string                     `json:"request_id"` // 请求 ID，标识一次请求
	ErrorCode int64                      `json:"error_code"` // openapi 错误码
	ErrorMsg  string                     `json:"error_msg"`  // openapi 错误信息
}

// GetCommentList
func GetCommentList(params *GetCommentListRequest) (*GetCommentListResponsedata, error) {
	var (
		err        error
		defaultRet *GetCommentListResponsedata
	)
	respData := &GetCommentListResponse{}

	client := NewHTTPClient().
		SetContentType(ContentTypeForm).
		SetConverterType(ConverterTypeJSON).
		SetMethod("POST").
		SetScheme(SCHEME).
		SetHost(OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/ma/component/comment/open_list")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("host_name", params.HostName)
	client.AddGetParam("sp_sdk_ver", SDKVERSION)
	client.AddGetParam("sp_sdk_lang", SDKLANG)
	client.AddPostParam("snid", params.Snid)
	client.AddPostParam("snid_type", params.SnidType)
	client.AddPostParam("start", fmt.Sprintf("%v", params.Start))
	client.AddPostParam("num", fmt.Sprintf("%v", params.Num))

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

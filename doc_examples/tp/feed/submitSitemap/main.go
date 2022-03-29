// 本示例基于百度智能小程序服务端开发者 smartapp-openapi-go SDK
// 使用该示例需要首先下载该 SDK，如果是第一次使用，可使用以下命令拉取依赖
// go get github.com/baidu/smartapp-openapi-go
// 如使用过程中遇到问题，可以加入如流群：5702992，进行反馈咨询
package main

import (
	"fmt"

	"github.com/baidu/smartapp-openapi-go/openapi/tp/feed"
	"github.com/baidu/smartapp-openapi-go/utils"
)

func main() {
	// 开发者在此设置请求参数，文档示例中的参数均为示例参数，实际参数请参考对应接口的文档上方的参数说明填写
	// 注意：代码示例中的参数字段基本是驼峰形式，而文档中的参数说明的参数字段基本是下划线形式
	// 如果开发者不想传非必需参数，可以将设置该参数的行注释
	reqParams := &feed.SubmitSitemapRequest{
		AccessToken: "#token",                // 文档中对应字段：access_token，实际使用时请替换成真实参数
		AppID:       158,                     // 文档中对应字段：app_id，实际使用时请替换成真实参数
		Desc:        "智能小程序示例",               // 文档中对应字段：desc，实际使用时请替换成真实参数
		Frequency:   3,                       // 文档中对应字段：frequency，实际使用时请替换成真实参数
		Type:        1,                       // 文档中对应字段：type，实际使用时请替换成真实参数
		URL:         "https://www.baidu.com", // 文档中对应字段：url，实际使用时请替换成真实参数
	}

	resp, err := feed.SubmitSitemap(reqParams)
	if err != nil {
		if _, ok := err.(*utils.OpenAPIError); ok {
			// openapi error
			// 可能是 access_token 无效，可以尝试重新生成 access_token
			fmt.Println("openapi error, ", err)
		} else if _, ok := err.(*utils.APIError); ok {
			// api error
			// 可能是参数错误或没有权限，建议根据错误信息自查或者社区发帖、加群反馈
			fmt.Println("api error, ", err)
		} else {
			// 其他错误
			fmt.Println("others error, ", err)
		}
	} else {
		fmt.Printf("%#v\n", resp)
	}
}

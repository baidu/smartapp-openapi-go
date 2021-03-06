// 本示例基于百度智能小程序服务端开发者 OpenAPI-SDK-Go
// 使用该示例需要首先下载该 SDK，使用引导见：https://smartprogram.baidu.com/docs/develop/serverapi/introduction_for_openapi_sdk/
// 使用之前请先确认下 SDK 版本是否为最新版本，如不是，请下载最新版本使用
// 如使用过程中遇到问题，可以加入如流群：5702992，进行反馈咨询
package main

import (
	"fmt"
	"os"

	"github.com/baidu/smartapp-openapi-go/openapi"
)

func main() {
	// 开发者在此设置请求参数，文档示例中的参数均为示例参数，实际参数请参考对应接口的文档上方的参数说明填写
	// 注意：代码示例中的参数字段基本是驼峰形式，而文档中的参数说明的参数字段基本是下划线形式
	// 如果开发者不想传非必需参数，可以将设置该参数的行注释
	fileReader, err := os.Open("xxxxx")
	if err != nil {
		panic(err)
	}
	defer fileReader.Close()

	checkImageRequestFileVal := openapi.CheckImageRequestFile{
		Name:   "image",
		Reader: fileReader,
	}

	reqParams := &openapi.CheckImageRequest{
		AccessToken: "24.134f41441a1883a5c8ad3c833652eb62.1211155.1563683264.710782-68858048", // 文档中对应字段：access_token，实际使用时请替换成真实参数
		Image:       checkImageRequestFileVal,                                                 // 文档中对应字段：image，实际使用时请替换成真实参数
		Type:        "porn,ocr-word,ocr-lead",                                                 // 文档中对应字段：type，实际使用时请替换成真实参数
	}

	resp, err := openapi.CheckImage(reqParams)
	if err != nil {
		if _, ok := err.(*openapi.OpenAPIError); ok {
			// openapi error
			// 可能是 access_token 无效，可以尝试重新生成 access_token
			fmt.Println("openapi error, ", err)
		} else if _, ok := err.(*openapi.APIError); ok {
			// api error
			// 可能是参数错误或没有权限，建议根据错误信息自查或者社区发帖、加群反馈
			fmt.Println("api error, ", err)
		} else {
			// 其他错误
			fmt.Println("others error, ", err)
		}
	} else {
		fmt.Printf("%#v\n", *resp)
	}
}

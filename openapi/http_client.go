package openapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

var debugFlag bool

const (
	ContentTypeForm      = "application/x-www-form-urlencoded"
	ContentTypeJSON      = "application/json"
	ContentTypeMultiPart = "multipart/form-data"
	ConverterTypeJSON    = "json"
	defaultTimeout       = 3 * time.Second
)

// httpClient 封装http请求
type httpClient struct {
	scheme        string
	host          string
	path          string
	method        string
	contentType   string
	converterType string
	config        *config
	getParams     url.Values
	postParams    url.Values
	requestBody   []byte
	headers       map[string]string
	rawResponse   []byte
	request       *http.Request
}

func init() {
	debugEnv := os.Getenv("DEBUG")
	debugL := strings.Split(debugEnv, ",")
	//如果环境变量包含swansdk字符串，开启debug日志
	for _, v := range debugL {
		if v == "swansdk" {
			debugFlag = true
			break
		}
	}
}

// newHTTPClient 创建一个HTTPClient
// opts 支持optTimeout 或 optRetry
func NewHTTPClient(opts ...Option) *httpClient {
	cfg := &config{
		retry:   0,
		timeout: defaultTimeout,
	}
	for _, opt := range opts {
		opt.apply(cfg)
	}
	return &httpClient{
		getParams:  url.Values{},
		postParams: url.Values{},
		headers:    map[string]string{},
		config:     cfg,
	}
}

func (hc *httpClient) SetContentType(contentType string) *httpClient {
	hc.contentType = contentType
	return hc
}
func (hc *httpClient) SetPath(path string) *httpClient {
	hc.path = path
	return hc
}
func (hc *httpClient) SetHost(host string) *httpClient {
	hc.host = host
	return hc
}

func (hc *httpClient) SetScheme(scheme string) *httpClient {
	hc.scheme = scheme
	return hc
}

func (hc *httpClient) SetMethod(method string) *httpClient {
	hc.method = method
	return hc
}

func (hc *httpClient) SetConverterType(converterType string) *httpClient {
	hc.converterType = converterType
	return hc
}

func (hc *httpClient) SetBody(input interface{}) *httpClient {
	switch input.(type) {
	case []byte:
		hc.requestBody = input.([]byte)
	case *(bytes.Buffer):
		tmp, _ := input.(*(bytes.Buffer))
		hc.requestBody = tmp.Bytes()
	default:
		bts, _ := json.Marshal(input)
		hc.requestBody = bts
	}
	return hc
}
func (hc *httpClient) AddPostParam(k, v string) *httpClient {
	hc.postParams.Add(k, v)
	return hc
}

func (hc *httpClient) AddGetParam(k, v string) *httpClient {
	hc.getParams.Add(k, v)
	return hc
}

func (hc *httpClient) AddHeader(k, v string) *httpClient {
	hc.headers[k] = v
	return hc
}

func (hc *httpClient) prepareRequest() error {
	reqURI := fmt.Sprintf("%s://%s%s", hc.scheme, hc.host, hc.path)
	if len(hc.getParams) > 0 {
		reqURI = fmt.Sprintf("%s?%s", reqURI, hc.getParams.Encode())
	}
	hc.debugLog("req_uri", reqURI)
	if hc.method == http.MethodGet {
		req, err := http.NewRequest(hc.method, reqURI, nil)
		if err != nil {
			hc.debugLog("getreq err %s", err)
			return err
		}
		hc.request = req
		return nil
	}
	var bodyReader io.Reader
	switch {
	case hc.contentType == ContentTypeForm:
		bodyReader = strings.NewReader(hc.postParams.Encode())
	case hc.contentType == ContentTypeJSON:
		bodyReader = strings.NewReader(string(hc.requestBody))
	case hc.contentType == ContentTypeMultiPart || strings.Contains(hc.contentType, ContentTypeMultiPart):
		bodyReader = bytes.NewBuffer(hc.requestBody)
	default:

	}

	req, err := http.NewRequest(hc.method, reqURI, bodyReader)
	if err != nil {
		hc.debugLog("postreq err %s", err)
		return err
	}
	req.Header.Add("content-type", hc.contentType)
	for k, v := range hc.headers {
		req.Header.Add(k, v)
	}
	hc.debugLog("http-req %#v", req)
	hc.request = req
	return nil
}

func (hc *httpClient) debugLog(format string, v ...interface{}) {
	if debugFlag {
		log.Printf(format, v...)
	}
}

func (hc *httpClient) Do() error {
	if err := hc.prepareRequest(); err != nil {
		return err
	}
	client := &http.Client{Timeout: hc.config.timeout}
	//todo retry && hook
	res, err := client.Do(hc.request)
	hc.debugLog("http response: %#v", res)
	if err != nil {
		return err
	}
	//错误码非20x
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return fmt.Errorf("status[%s]", res.Status)
	}
	if res.Body == nil {
		return fmt.Errorf("nil body")
	}
	defer res.Body.Close()
	btsRes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	hc.rawResponse = btsRes
	hc.debugLog("raw res: %s", btsRes)
	return nil
}

func (hc *httpClient) GetRawResponse() []byte {
	return hc.rawResponse
}
func (hc *httpClient) Convert(resp interface{}) error {
	switch hc.converterType {
	case ConverterTypeJSON:
		return json.Unmarshal(hc.rawResponse, resp)
	default:
		return fmt.Errorf("invalid converter[%s]", hc.converterType)
	}
}

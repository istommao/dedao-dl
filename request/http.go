package request

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strconv"
	"strings"
	"time"

	jsoniter "github.com/json-iterator/go"
)

var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary
	// UserAgent UserAgent
	UserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_1_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36"
)

// var client *http.Client
const (
	MaxIdleConns        = 100
	MaxIdleConnsPerHost = 100
	IdleConnTimeout     = 90
	Timeout             = 30
)

// HTTPClient http client
type HTTPClient struct {
	BaseURL string
	Header  map[string]string
	Data    interface{}
	Client  *http.Client
}

// NewClient new HTTPClient
func NewClient(baseURL string) *HTTPClient {
	c := &HTTPClient{
		BaseURL: baseURL,
		Header:  make(map[string]string),
		Data:    make(map[string]string),
		Client:  createHTTPClient(),
	}
	c.ResetCookieJar()
	return c
}

// createHTTPClient for connection re-use
func createHTTPClient() *http.Client {
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   Timeout * time.Second,
				KeepAlive: Timeout * time.Second,
			}).DialContext,
			MaxIdleConns:        MaxIdleConns,
			MaxIdleConnsPerHost: MaxIdleConnsPerHost,
			IdleConnTimeout:     IdleConnTimeout * time.Second,
		},
	}
	return client
}

//SetCookiejar set Cookie
func (h *HTTPClient) SetCookiejar(jar http.CookieJar) *HTTPClient {
	h.Client.Jar = jar
	return h
}

//ResetCookieJar reset cookie jar
func (h *HTTPClient) ResetCookieJar() *HTTPClient {
	h.Client.Jar, _ = cookiejar.New(nil)
	return h
}

//SetTimeout set timeout
func (h *HTTPClient) SetTimeout(t time.Duration) *HTTPClient {
	h.Client.Timeout = t
	return h
}

// SetBaseURL 设置基础地址
func (h *HTTPClient) SetBaseURL(baseURL string) *HTTPClient {
	h.BaseURL = baseURL
	return h
}

// AddHeader 添加一个 header
func (h *HTTPClient) AddHeader(name string, value string) *HTTPClient {
	h.Header[name] = value
	return h
}

// AddHeaders 添加一组 header 数据
func (h *HTTPClient) AddHeaders(headers map[string]string) *HTTPClient {
	for k, header := range headers {
		h.Header[k] = header
	}
	return h
}

// SetData 设置查询数据
func (h *HTTPClient) SetData(data interface{}) *HTTPClient {
	h.Data = data
	return h
}

// Request 实现 http／https 访问，
// 根据给定的 method (GET, POST, HEAD, PUT 等等),
// URL (路由),
// 返回值分别为 *http.Response, 错误信息
func (h *HTTPClient) Request(method, URL string) (*http.Response, error) {
	var (
		req   *http.Request
		obody io.Reader
	)
	requrl := h.BaseURL + URL
	if h.Data != nil {
		switch value := h.Data.(type) {
		case io.Reader:
			obody = value
		case map[string]string, map[string]int, map[string]interface{}, []int, []string:
			if method == "GET" {
				query := url.Values{}
				if params, ok := h.Data.(map[string]string); ok {
					for k, v := range params {
						query.Add(k, v)
					}
				}
				encode := query.Encode()
				if encode != "" {
					requrl += "?" + encode
				}
			} else {
				postData, err := jsoniter.Marshal(value)
				if err != nil {
					return nil, err
				}
				h.Header["Content-Type"] = "application/json"
				obody = bytes.NewReader(postData)
			}
		case string:
			obody = strings.NewReader(value)
		case []byte:
			obody = bytes.NewReader(value)
		default:
			return nil, fmt.Errorf("request.Req: unknow post type: %s", h.Data)
		}
	}
	// fmt.Println(method, " ", requrl)
	req, err := http.NewRequest(method, requrl, obody)
	if err != nil {
		return nil, err
	}

	// 设置浏览器标识
	req.Header.Set("User-Agent", UserAgent)

	if h.Header != nil {
		for k, v := range h.Header {
			req.Header.Set(k, v)
		}
	}
	return h.Client.Do(req)
}

// HTTPGet http get request
func HTTPGet(url string) (body []byte, err error) {
	r, err := Get(url)
	defer r.Close()
	if err != nil {

	}
	body, err = ioutil.ReadAll(r)
	if err != nil {
		return
	}
	return
}

// Get http get request
func Get(url string) (io.ReadCloser, error) {
	client := NewClient(url)
	resp, err := client.Request("GET", "")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http error: status code %d", resp.StatusCode)
	}
	return resp.Body, nil
}

// Headers return the HTTP Headers of the url
func Headers(url string) (http.Header, error) {
	client := NewClient(url)
	resp, err := client.Request("GET", "")
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	return resp.Header, nil
}

// Size get size of the url
func Size(url string) (int, error) {
	h, err := Headers(url)
	if err != nil {
		return 0, err
	}
	s := h.Get("Content-Length")
	if s == "" {
		return 0, errors.New("Content-Length is not present")
	}
	size, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return size, nil
}

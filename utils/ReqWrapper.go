// utils/ReqWrapper.go

package utils

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

func newTanRequest(method string, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("accept-language", "zh-CN, zh, en-US;q=0.9, en;q=0.8")
	req.Header.Set("sec-ch-ua", "\"Not(A:Brand\";v=\"99\", \"Brave\";v=\"133\", \"Chromium\";v=\"133\"")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", "\"Windows\"")
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("sec-gpc", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.1; x64; zh-CN) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/133.0.0.0 Safari/537.36")

	return req, nil
}

func AddHeader(req *http.Request, key string, value string) {
	req.Header.Set(key, value)
}

func JsonString2Body(body string) io.Reader {
	return io.NopCloser(strings.NewReader(body))
}

func Object2Body(body interface{}) io.Reader {
	str, err := json.Marshal(body)
	if err != nil {
		return nil
	}
	return JsonString2Body(string(str))
}

func NewTanGetRequest(url string) (*http.Request, error) {
	return newTanRequest("GET", url, nil)
}

func NewTanPostRequest(url string, body io.Reader) (*http.Request, error) {
	return newTanRequest("POST", url, body)
}

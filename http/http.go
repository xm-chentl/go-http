package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	gohttp "github.com/xm-chentl/go-http"
)

// New net/http实现，是个常规使用方式
func New() gohttp.IHttp {
	return &gohttp.BaseHTTP{
		HandleFunc: func(method, url string, queryArgs, requestData, responseData interface{}, headers map[string]string) (respData gohttp.HttpResponse, err error) {
			var bodyBytes []byte
			switch requestData := requestData.(type) {
			case string:
				bodyBytes = []byte(requestData)
			case []byte:
				bodyBytes = requestData
			default:
				bodyBytes, err = json.Marshal(requestData)
				if err != nil {
					err = fmt.Errorf("requestData to byte[] is fail, err(%v)", err)
					return
				}
			}
			if queryArgs != nil {
				var queryArgsBytes []byte
				queryArgsBytes, err = toValues(queryArgs)
				if err != nil {
					return
				}
				url += "?" + string(queryArgsBytes)
			}

			req, err := http.NewRequest(method, url, bytes.NewBuffer(bodyBytes))
			if err != nil {
				return
			}

			req.Header.Add("Content-Type", gohttp.ContentType)
			for k, v := range headers {
				req.Header.Add(k, v)
			}

			client := &http.Client{Timeout: 30 * time.Second}
			resp, err := client.Do(req)
			if err != nil {
				return
			}

			respData = gohttp.HttpResponse{}
			defer resp.Body.Close()
			respData.Body, err = ioutil.ReadAll(resp.Body)
			if err != nil {
				return
			}

			respData.Code = resp.StatusCode
			if responseData != nil {
				// 如存在返回结构则反序列化
				if err = json.Unmarshal(respData.Body, responseData); err != nil {
					respData.Error = err
				}
			}

			return
		},
	}
}

func toValues(requestArgs interface{}) (res []byte, err error) {
	argsBytes, err := json.Marshal(requestArgs)
	if err != nil {
		return
	}

	argsMap := make(map[string]interface{})
	if err = json.Unmarshal(argsBytes, &argsMap); err != nil {
		return
	}

	argsArray := make([]string, 0)
	for k, v := range argsMap {
		argsArray = append(argsArray, fmt.Sprintf("%s=%s", k, v))

	}
	res = []byte(strings.Join(argsArray, "&"))

	return
}

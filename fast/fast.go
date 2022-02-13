package fast

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/valyala/fasthttp"

	gohttp "github.com/xm-chentl/go-http"
)

// New fasthttp实现，是个具体强规范的请求组件
func New() gohttp.IHttp {
	return &gohttp.BaseHTTP{
		HandleFunc: func(method, url string, queryArgs, requestData, responseData interface{}, headers map[string]string) (respData gohttp.HttpResponse, err error) {
			req := fasthttp.AcquireRequest()
			defer fasthttp.ReleaseRequest(req)

			req.Header.SetContentType(gohttp.ContentType)
			req.Header.SetMethod(method)

			if queryArgs != nil {
				var queryArgsBytes []byte
				if queryArgsBytes, err = toValues(queryArgs); err != nil {
					return
				}
				if len(queryArgsBytes) > 0 {
					url += "?" + string(queryArgsBytes)
				}
			}
			req.SetRequestURI(url)
			if method == http.MethodPost {
				var requestDataOfByte []byte
				requestDataOfByte, err = json.Marshal(requestData)
				if err != nil {
					err = fmt.Errorf("fasthttp requestData to []byte fail err: %v", err)
					return
				}
				req.SetBody(requestDataOfByte)
			}

			for k, v := range headers {
				req.Header.Add(k, v)
			}

			resp := fasthttp.AcquireResponse()
			defer fasthttp.ReleaseResponse(resp)

			if err = fasthttp.Do(req, resp); err != nil {
				err = fmt.Errorf("fasthttp %s request fail err: %v", string(req.Header.Method()), err)
				return
			}

			respData = gohttp.HttpResponse{}
			if responseData != nil {
				if err = json.Unmarshal(resp.Body(), responseData); err != nil {
					err = fmt.Errorf("[]byte to responseData fail, err: %v", err)
					return
				}

				return
			}
			respData.Code = resp.StatusCode()
			respData.Body = resp.Body()

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

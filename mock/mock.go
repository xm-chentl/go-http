package mock

import (
	"encoding/json"

	gohttp "github.com/xm-chentl/go-http"
)

// New 实例一个mock实例
func New(callback func(method, url string, requestData interface{}, header map[string]string) (interface{}, error)) gohttp.IHttp {
	return &gohttp.BaseHTTP{
		HandleFunc: func(method, url string, queryArgs, requestData, responseData interface{}, header map[string]string) (resp gohttp.HttpResponse, err error) {
			respData, err := callback(method, url, requestData, header)
			if err != nil {
				return
			}

			// 序列化，使用mock者无需要关注
			respDataByte, _ := json.Marshal(respData)
			err = json.Unmarshal(respDataByte, responseData)
			if err != nil {
				return
			}

			return
		},
	}
}

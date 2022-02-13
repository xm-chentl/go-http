package gohttp

import (
	"fmt"
	"net/http"
	"reflect"
)

// HandleFunc 处理函数
type HandleFunc func(method, url string, queryArgs, requestData, responseData interface{}, headers map[string]string) (HttpResponse, error)

// BaseHTTP 基础请求类
type BaseHTTP struct {
	url         string
	method      string
	queryArgs   interface{}
	requestData interface{}
	headers     map[string]string
	HandleFunc  HandleFunc
}

// SetURL 设置请求地址
func (b *BaseHTTP) SetURL(url string) IHttp {
	b.url = url
	return b
}

func (b *BaseHTTP) SetQueryArgs(queryArgs interface{}) IHttp {
	b.queryArgs = queryArgs
	return b
}

// SetBody 设置请求的参数
func (b *BaseHTTP) SetBody(requestData interface{}) IHttp {
	b.requestData = requestData
	return b
}

// SetHeader 设置表头
func (b *BaseHTTP) SetHeader(headers map[string]string) IHttp {
	b.headers = headers
	return b
}

// SetMethod 设置请求方式
func (b *BaseHTTP) SetMethod(method string) IHttp {
	b.method = method
	return b
}

// Send 发送请求
func (b *BaseHTTP) Send(responseData interface{}) (resp HttpResponse, err error) {
	defer b.Reset()
	if responseData != nil && reflect.TypeOf(responseData).Kind() != reflect.Ptr {
		err = fmt.Errorf("receive parameter responseData must ptr")
		return
	}
	if b.method == "" {
		// 默认post
		b.method = http.MethodPost
	}

	return b.HandleFunc(b.method, b.url, b.queryArgs, b.requestData, responseData, b.headers)
}

// Reset 重置参数
func (b *BaseHTTP) Reset() {
	b.url = ""
	b.method = ""
	b.queryArgs = nil
	b.requestData = nil
	b.headers = make(map[string]string)
}

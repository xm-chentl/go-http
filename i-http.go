package gohttp

// IHttp http 统一接口
type IHttp interface {
	SetMethod(string) IHttp
	SetURL(string) IHttp
	SetQueryArgs(interface{}) IHttp
	SetBody(interface{}) IHttp
	SetHeader(map[string]string) IHttp
	Send(interface{}) (resp HttpResponse, err error)
}
